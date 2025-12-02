package helper

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"maps"

	"swiftschool/config"
	"swiftschool/domain"

	"github.com/hashicorp/vault/api"
)

// -------------------------
// Vault Client Helper
// -------------------------
func NewVaultClient(cfg config.AppConfig) (*api.Client, error) {
	logger := GetLogger()

	client, err := api.NewClient(&api.Config{Address: cfg.VaultAddr})
	if err != nil {
		logger.Errorf("vault client creation failed: %v", err)
		return nil, err
	}

	client.SetToken(cfg.VaultToken)
	return client, nil
}

// -------------------------
// KV Version Detection
// -------------------------
func getKVVersion(client *api.Client, enginePath string) (int, error) {
	logger := GetLogger()

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		logger.Errorf("list mounts failed: %v", err)
		return 0, err
	}

	mount, ok := mounts[enginePath+"/"]
	if !ok {
		logger.Errorf("mount path not found: %s", enginePath)
		return 0, errors.New("mount path not found")
	}

	if mount.Options["version"] == "2" {
		return 2, nil
	}

	return 1, nil
}

// -------------------------
// Key Management
// -------------------------
func GetInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	logger := GetLogger()
	client, err := NewVaultClient(cfg)
	if err != nil {
		return nil, err
	}

	enginePath := cfg.KVEnginePath
	if enginePath == "" {
		enginePath = domain.KVEnginePath
	}

	kvVersion, err := getKVVersion(client, enginePath)
	if err != nil {
		logger.Errorf("KV version detection failed: %v", err)
		return nil, err
	}

	var secretData map[string]any

	if kvVersion == 2 {
		secret, err := client.KVv2(enginePath).Get(ctx, domain.InstituteSecretKey)
		if err != nil {
			logger.Errorf("failed to fetch KVv2 secret: %v", err)
			return nil, err
		}
		if secret == nil || secret.Data == nil {
			logger.Errorf("secret %s missing or empty", domain.InstituteSecretKey)
			return nil, errors.New("secret missing")
		}
		secretData = secret.Data

	} else { // KV v1
		secret, err := client.Logical().Read(enginePath + "/" + domain.InstituteSecretKey)
		if err != nil {
			logger.Errorf("failed to fetch KVv1 secret: %v", err)
			return nil, err
		}
		if secret == nil || secret.Data == nil {
			logger.Errorf("KVv1 secret empty: %s", domain.InstituteSecretKey)
			return nil, errors.New("secret missing")
		}
		secretData = secret.Data
	}

	val, ok := secretData[instituteID].(string)
	if !ok || val == "" {
		logger.Errorf("institute key not found: %s", instituteID)
		return nil, errors.New("institute key not found")
	}

	key, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		logger.Errorf("failed to base64 decode key: %v", err)
		return nil, err
	}

	keyLen := cfg.AESKeyLength
	if keyLen == 0 {
		keyLen = domain.DefaultAESKeyLen
	}

	if len(key) != keyLen {
		logger.Errorf("invalid AES key length. expected=%d actual=%d", keyLen, len(key))
		return nil, domain.ErrInvalidKeyLength
	}

	return key, nil
}

func EnsureInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	logger := GetLogger()

	key, err := GetInstituteKey(ctx, cfg, instituteID)
	if err == nil {
		return key, nil
	}

	logger.Errorf("key missing, generating new one for institute %s", instituteID)

	keyStr, err := CreateInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		logger.Errorf("failed to create institute key: %v", err)
		return nil, err
	}

	key, err = base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		logger.Errorf("failed to decode newly created key: %v", err)
		return nil, err
	}

	return key, nil
}

func CreateInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) (string, error) {
	logger := GetLogger()

	client, err := NewVaultClient(cfg)
	if err != nil {
		return "", err
	}

	keyLen := cfg.AESKeyLength
	if keyLen == 0 {
		keyLen = domain.DefaultAESKeyLen
	}

	keyBytes := make([]byte, keyLen)
	if _, err := io.ReadFull(rand.Reader, keyBytes); err != nil {
		logger.Errorf("AES key generation failed: %v", err)
		return "", err
	}

	keyBase64 := base64.StdEncoding.EncodeToString(keyBytes)

	enginePath := cfg.KVEnginePath
	if enginePath == "" {
		enginePath = domain.KVEnginePath
	}

	kvVersion, err := getKVVersion(client, enginePath)
	if err != nil {
		logger.Errorf("KV version detection failed: %v", err)
		return "", err
	}

	var data = map[string]any{}

	if kvVersion == 2 {
		secret, _ := client.KVv2(enginePath).Get(ctx, domain.InstituteSecretKey)
		if secret != nil && secret.Data != nil {
			maps.Copy(data, secret.Data)

		}

		data[instituteID] = keyBase64

		_, err := client.KVv2(enginePath).Put(ctx, domain.InstituteSecretKey, data)
		if err != nil {
			logger.Errorf("storing KVv2 key failed: %v", err)
			return "", err
		}

	} else { // KV1
		secret, _ := client.Logical().Read(enginePath + "/" + domain.InstituteSecretKey)
		if secret != nil && secret.Data != nil {
			maps.Copy(data, secret.Data)

		}

		data[instituteID] = keyBase64

		_, err := client.Logical().Write(enginePath+"/"+domain.InstituteSecretKey, data)
		if err != nil {
			logger.Errorf("storing KVv1 key failed: %v", err)
			return "", err
		}
	}

	return keyBase64, nil
}

// -------------------------
// AES-GCM Encryption / Decryption
// -------------------------
func EncryptText(ctx context.Context, cfg config.AppConfig, instituteID, plaintext string) (string, error) {
	logger := GetLogger()

	key, err := EnsureInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		logger.Errorf("failed to fetch key for encryption: %v", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("AES cipher creation failed: %v", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("GCM creation failed: %v", err)
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Errorf("nonce creation failed: %v", err)
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptText(ctx context.Context, cfg config.AppConfig, instituteID, cipherTextB64 string) (string, error) {
	logger := GetLogger()

	key, err := EnsureInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		logger.Errorf("failed to fetch key for decryption: %v", err)
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		logger.Errorf("ciphertext base64 decode failed: %v", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("AES cipher creation failed: %v", err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("GCM creation failed: %v", err)
		return "", err
	}

	if len(cipherText) < aesgcm.NonceSize() {
		logger.Errorf("ciphertext too short")
		return "", errors.New("ciphertext too short")
	}

	nonce := cipherText[:aesgcm.NonceSize()]
	encrypted := cipherText[aesgcm.NonceSize():]

	plaintext, err := aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		logger.Errorf("decryption failed: %v", err)
		return "", err
	}

	return string(plaintext), nil
}
