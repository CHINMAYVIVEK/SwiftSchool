package config

import (
	"time"
)

// AppConfig defines the application's configuration structure.
// It maps environment variables to struct fields using tags.
type AppConfig struct {
	// Server & Logging
	LogLevel     string        `env:"LOG_LEVEL" default:"debug"`
	ServerPort   string        `env:"SERVER_PORT" default:":8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"15s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"15s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" default:"60s"`

	// Vault configuration
	VaultAddr    string `env:"VAULT_ADDR"`           // Vault server address (e.g., http://localhost:8200)
	VaultToken   string `env:"VAULT_TOKEN"`          // Vault authentication token
	VaultCACert  string `env:"VAULT_CA_CERT"`        // Path to CA Certificate for TLS (Added for Prod)
	KVEnginePath string `env:"VAULT_KV_ENGINE_PATH"` // Vault KV secret engine mount path

	// Cryptography
	AESKeyLength int `env:"AES_KEY_LENGTH" default:"32"` // AES-256 key length (32 bytes)
}
