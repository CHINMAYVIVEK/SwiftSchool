package config

import "time"

// AppConfig holds general application configuration
type AppConfig struct {
	// Server & Logging
	LogLevel     string        `env:"LOG_LEVEL" default:"debug"`
	ServerPort   string        `env:"SERVER_PORT" default:":8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"15s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"15s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" default:"60s"`

	// Vault configuration
	VaultAddr    string `env:"VAULT_ADDR"`           // Vault server address
	VaultToken   string `env:"VAULT_TOKEN"`          // Vault token
	KVEnginePath string `env:"VAULT_KV_ENGINE_PATH"` // Vault KV path
	InstituteID  string `env:"INSTITUTE_ID"`         // Current institute ID

	// Cryptography
	AESKeyLength int `env:"AES_KEY_LENGTH" default:"32"` // AES-256 key length

}
