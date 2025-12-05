package helper

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"maps"
	"sync"

	"swiftschool/config"
	"swiftschool/domain"

	"github.com/hashicorp/vault/api"
)

// =================================================================================
// GLOBAL KEY CACHE (PERFORMANCE OPTIMIZATION)
// =================================================================================
// In a high-throughput system, we cannot hit Vault for every DB row.
// We cache keys in memory. In a distributed system, key rotation events
// would typically propagate via a message bus (Redis/NATS), but for this
// modular monolith, a package-level sync.Map is highly efficient.
var (
	keyCache sync.Map // map[string][]byte (InstituteID -> Raw Key)
)

// InvalidateKeyCache clears a specific key from memory, forcing a fresh fetch next time.
func InvalidateKeyCache(instituteID string) {
	keyCache.Delete(instituteID)
}

// =================================================================================
// SECTION 1: VAULT INFRASTRUCTURE & CLIENT
// =================================================================================

// NewVaultClient establishes a secure connection to the Vault server.
func NewVaultClient(cfg config.AppConfig) (*api.Client, error) {
	logger := GetLogger()

	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = cfg.VaultAddr

	if cfg.VaultCACert != "" {
		if err := vaultConfig.ConfigureTLS(&api.TLSConfig{CACert: cfg.VaultCACert}); err != nil {
			logger.Errorf("vault: failed to configure TLS: %v", err)
			return nil, fmt.Errorf("%w: TLS configuration failed", domain.ErrVaultConnection)
		}
	}

	client, err := api.NewClient(vaultConfig)
	if err != nil {
		logger.Errorf("vault: client creation failed: %v", err)
		return nil, domain.ErrVaultConnection
	}

	if cfg.VaultToken != "" {
		client.SetToken(cfg.VaultToken)
	}

	return client, nil
}

// DetectKVVersion determines if the engine is KV v1 or v2.
func DetectKVVersion(client *api.Client, enginePath string) (domain.KVVersion, error) {
	logger := GetLogger()

	if enginePath == "" {
		enginePath = domain.KVEnginePath
	}
	mountPath := enginePath + "/"

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		logger.Errorf("vault: failed to list mounts: %v", err)
		return 0, domain.ErrVaultConnection
	}

	mount, ok := mounts[mountPath]
	if !ok {
		// Optimization: Assume KV2 if strict match fails to avoid hard errors in some setups
		return domain.KVVersion2, nil
	}

	if mount.Options[domain.KvVersionKey] == "2" {
		return domain.KVVersion2, nil
	}

	return domain.KVVersion1, nil
}

// =================================================================================
// SECTION 2: KEY LIFECYCLE MANAGEMENT (CRUD)
// =================================================================================

// GetInstituteKey retrieves the raw AES-256 key.
// OPTIMIZATION: Checks memory cache first. Hits Vault only on cache miss.
func GetInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	// 1. Fast Path: Check Cache
	if cached, ok := keyCache.Load(instituteID); ok {
		return cached.([]byte), nil
	}

	// 2. Slow Path: Fetch from Vault
	logger := GetLogger()
	client, err := NewVaultClient(cfg)
	if err != nil {
		return nil, err
	}

	enginePath := cfg.KVEnginePath
	if enginePath == "" {
		enginePath = domain.KVEnginePath
	}

	version, err := DetectKVVersion(client, enginePath)
	if err != nil {
		return nil, err
	}

	var secretData map[string]interface{}

	if version == domain.KVVersion2 {
		secret, err := client.KVv2(enginePath).Get(ctx, domain.InstituteSecretKey)
		if err != nil {
			if errors.Is(err, api.ErrSecretNotFound) {
				return nil, domain.ErrSecretNotFound
			}
			logger.Errorf("vault: KVv2 get failed: %v", err)
			return nil, err
		}
		if secret == nil || secret.Data == nil {
			return nil, domain.ErrSecretNotFound
		}
		secretData = secret.Data
	} else {
		path := fmt.Sprintf("%s/%s", enginePath, domain.InstituteSecretKey)
		secret, err := client.Logical().Read(path)
		if err != nil {
			logger.Errorf("vault: KVv1 read failed: %v", err)
			return nil, err
		}
		if secret == nil || secret.Data == nil {
			return nil, domain.ErrSecretNotFound
		}
		secretData = secret.Data
	}

	val, ok := secretData[instituteID]
	if !ok {
		return nil, domain.ErrKeyNotFound
	}

	keyStr, ok := val.(string)
	if !ok || keyStr == "" {
		return nil, domain.ErrKeyNotFound
	}

	keyBytes, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		logger.Errorf("vault: corrupt key data for %s: %v", instituteID, err)
		return nil, fmt.Errorf("integrity check failed")
	}

	// Use Configured Key Length (defaulting to Domain constant if 0)
	expectedLen := cfg.AESKeyLength
	if expectedLen == 0 {
		expectedLen = domain.DefaultAESKeyLen
	}

	if len(keyBytes) != expectedLen {
		logger.Errorf("vault: invalid key length for %s: got %d expected %d", instituteID, len(keyBytes), expectedLen)
		return nil, domain.ErrInvalidKeyLength
	}

	// 3. Populate Cache
	keyCache.Store(instituteID, keyBytes)

	return keyBytes, nil
}

// EnsureInstituteKey guarantees a key exists. Uses Cache -> Vault -> Create flow.
func EnsureInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	key, err := GetInstituteKey(ctx, cfg, instituteID)
	if err == nil {
		return key, nil
	}

	if errors.Is(err, domain.ErrKeyNotFound) || errors.Is(err, domain.ErrSecretNotFound) {
		GetLogger().Infof("vault: Initializing NEW encryption key for institute: %s", instituteID)
		return RotateInstituteKey(ctx, cfg, instituteID)
	}

	return nil, err
}

// RotateInstituteKey generates a new key, saves to Vault, and updates Cache.
func RotateInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	logger := GetLogger()

	// 1. Generate Key using Configured Length
	keyLen := cfg.AESKeyLength
	if keyLen == 0 {
		keyLen = domain.DefaultAESKeyLen
	}

	newKey := make([]byte, keyLen)
	if _, err := io.ReadFull(rand.Reader, newKey); err != nil {
		return nil, errors.New("entropy failure")
	}
	keyBase64 := base64.StdEncoding.EncodeToString(newKey)

	// 2. Vault Connect
	client, err := NewVaultClient(cfg)
	if err != nil {
		return nil, err
	}

	enginePath := cfg.KVEnginePath
	if enginePath == "" {
		enginePath = domain.KVEnginePath
	}

	version, err := DetectKVVersion(client, enginePath)
	if err != nil {
		return nil, err
	}

	// 3. Optimistic Locking Read
	currentData := make(map[string]interface{})
	if version == domain.KVVersion2 {
		secret, _ := client.KVv2(enginePath).Get(ctx, domain.InstituteSecretKey)
		if secret != nil && secret.Data != nil {
			maps.Copy(currentData, secret.Data)
		}
	} else {
		path := fmt.Sprintf("%s/%s", enginePath, domain.InstituteSecretKey)
		secret, _ := client.Logical().Read(path)
		if secret != nil && secret.Data != nil {
			maps.Copy(currentData, secret.Data)
		}
	}

	// 4. Update
	currentData[instituteID] = keyBase64

	// 5. Write
	if version == domain.KVVersion2 {
		_, err = client.KVv2(enginePath).Put(ctx, domain.InstituteSecretKey, currentData)
	} else {
		path := fmt.Sprintf("%s/%s", enginePath, domain.InstituteSecretKey)
		_, err = client.Logical().Write(path, currentData)
	}

	if err != nil {
		logger.Errorf("vault: write failed for %s: %v", instituteID, err)
		return nil, fmt.Errorf("%w: persistence failed", domain.ErrVaultConnection)
	}

	// 6. Update Cache
	keyCache.Store(instituteID, newKey)
	logger.Infof("vault: successfully rotated and cached key for %s", instituteID)

	return newKey, nil
}

// =================================================================================
// SECTION 3: ON-DEMAND CRYPTOGRAPHY HELPERS
// =================================================================================

// Encrypt encrypts plain text string -> Base64 Ciphertext.
// Use this in your Service Layer BEFORE inserting into DB.
func Encrypt(ctx context.Context, cfg config.AppConfig, instituteID, plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	key, err := EnsureInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("crypto: cipher init failed: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("crypto: gcm init failed: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("crypto: nonce gen failed: %w", err)
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts Base64 Ciphertext -> Plain text string.
// Use this in your Service Layer AFTER fetching from DB.
func Decrypt(ctx context.Context, cfg config.AppConfig, instituteID, cipherTextB64 string) (string, error) {
	if cipherTextB64 == "" {
		return "", nil
	}

	data, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", errors.New("malformed ciphertext")
	}

	key, err := GetInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		GetLogger().Errorf("decrypt: auth failed for %s", instituteID)
		return "", errors.New("decryption failed: invalid auth tag")
	}

	return string(plaintext), nil
}
