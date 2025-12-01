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

	"swiftschool/config"

	"github.com/hashicorp/vault/api"
)

var ErrInvalidKeyLength = errors.New("encryption key must be 32 bytes")

// generateRandomKey generates a 32-byte AES key and returns it as Base64 string
func generateRandomKey() (string, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", fmt.Errorf("failed to generate random key: %w", err)
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

// CreateInstituteKeyInVault stores a new key for an institute in Vault (one-time call)
// The key will be stored in Vault as: { "<instituteID>": "<base64_key>" }
func CreateInstituteKeyInVault(ctx context.Context, cfg config.AppConfig, instituteID string) (string, error) {
	client, err := api.NewClient(&api.Config{Address: cfg.VaultAddr})
	if err != nil {
		return "", fmt.Errorf("failed to create Vault client: %w", err)
	}
	client.SetToken(cfg.VaultToken)

	// Check if key already exists
	_, err = client.KVv2(cfg.KVEnginePath).Get(ctx, instituteID)
	if err == nil {
		return "", fmt.Errorf("key already exists for institute %s", instituteID)
	}

	key, err := generateRandomKey()
	if err != nil {
		return "", err
	}

	// Store the key with instituteID as the field name
	data := map[string]interface{}{
		instituteID: key,
	}

	if _, err := client.KVv2(cfg.KVEnginePath).Put(ctx, instituteID, data); err != nil {
		return "", fmt.Errorf("failed to store key in Vault: %w", err)
	}

	return key, nil
}

// getInstituteKey fetches AES key for an institute from Vault (private)
func getInstituteKey(ctx context.Context, cfg config.AppConfig, instituteID string) ([]byte, error) {
	client, err := api.NewClient(&api.Config{Address: cfg.VaultAddr})
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}
	client.SetToken(cfg.VaultToken)

	secret, err := client.KVv2(cfg.KVEnginePath).Get(ctx, instituteID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch key from Vault for institute %s: %w", instituteID, err)
	}

	// Use instituteID as the field to get the key
	keyBase64, ok := secret.Data[instituteID].(string)
	if !ok {
		return nil, errors.New("key not found or invalid type in Vault")
	}

	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode key from Base64: %w", err)
	}

	if len(key) != cfg.AESKeyLength {
		return nil, ErrInvalidKeyLength
	}

	return key, nil
}

// EncryptAESGCM encrypts plaintext using AES-GCM, fetching the key from Vault
func EncryptAESGCM(ctx context.Context, cfg config.AppConfig, instituteID, plaintext string) (string, error) {
	key, err := getInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		return "", err
	}
	fmt.Println("Key ==>> ", base64.StdEncoding.EncodeToString(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAESGCM decrypts ciphertext using AES-GCM, fetching the key from Vault
func DecryptAESGCM(ctx context.Context, cfg config.AppConfig, instituteID, cipherTextB64 string) (string, error) {
	key, err := getInstituteKey(ctx, cfg, instituteID)
	if err != nil {
		return "", err
	}
	fmt.Println("Key ==>> ", base64.StdEncoding.EncodeToString(key))

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", fmt.Errorf("failed to decode Base64 ciphertext: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce := cipherText[:nonceSize]
	encrypted := cipherText[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt ciphertext: %w", err)
	}

	return string(plaintext), nil
}
