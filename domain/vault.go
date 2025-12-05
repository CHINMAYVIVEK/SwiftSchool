package domain

import (
	"errors"
	"fmt"
	"time"
)

// -------------------------
// Vault / KV Configuration
// -------------------------

var (
	// KVEnginePath is the mount path for the KV secret engine.
	// Can be overwritten by configuration at runtime.
	KVEnginePath = "secret"

	// InstituteSecretKey is the path/name of the secret holding all institute keys.
	// Structure: secret/institute-keys -> { "institute_id_1": "base64_key", "institute_id_2": "base64_key" }
	InstituteSecretKey = "institute-keys"

	// DefaultAESKeyLen defines the required length for AES-256 keys (32 bytes).
	DefaultAESKeyLen = 32
)

// -------------------------
// Vault Diagnostic Constants
// -------------------------
const (
	KvType       = "kv"        // KV type for secret engine
	KvVersionKey = "version"   // version key in mount options
	KvV2Metadata = "metadata/" // path prefix for KV v2 metadata
	KvV2Data     = "data/"     // path prefix for KV v2 data
)

// -------------------------
// Types & Enums
// -------------------------

// KVVersion represents the version of the Key-Value engine (v1 or v2).
type KVVersion int

const (
	KVVersion1 KVVersion = 1
	KVVersion2 KVVersion = 2
)

// AuthMethod defines the authentication method used for Vault.
type AuthMethod string

const (
	AuthToken      AuthMethod = "token"
	AuthAppRole    AuthMethod = "approle"
	AuthKubernetes AuthMethod = "kubernetes"
)

// -------------------------
// Errors
// -------------------------
var (
	// ErrInvalidKeyLength indicates the retrieved or generated key does not match the required AES length.
	ErrInvalidKeyLength = fmt.Errorf("encryption key must be %d bytes", DefaultAESKeyLen)

	// ErrSecretNotFound indicates the specific path or secret does not exist in Vault.
	ErrSecretNotFound = errors.New("vault secret not found")

	// ErrKeyNotFound indicates the secret exists, but the specific key for the institute is missing.
	ErrKeyNotFound = errors.New("institute key not found in secret")

	// ErrVaultConnection indicates a failure to communicate with the Vault server.
	ErrVaultConnection = errors.New("failed to connect to vault")
)

// -------------------------
// Structs
// -------------------------

// VaultSecret represents a generic secret stored in Vault.
// This mirrors the standard Vault response structure.
type VaultSecret struct {
	Path      string                 `json:"path"`
	Data      map[string]interface{} `json:"data"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Version   int                    `json:"version,omitempty"`
	CreatedAt time.Time              `json:"created_at,omitempty"`
}

// InstituteKey represents the internal domain model for an encryption key.
// Note: In the current Vault schema (InstituteSecretKey), keys are stored as a map of ID->KeyString.
// This struct is used when parsing or managing keys in the application layer.
type InstituteKey struct {
	InstituteID string    `json:"institute_id"`
	Key         string    `json:"key"` // Base64 encoded AES key
	CreatedAt   time.Time `json:"created_at"`
	Version     int       `json:"version"`
}

// VaultConfig holds comprehensive configuration for the Vault client.
type VaultConfig struct {
	Address    string     `json:"address"`
	Token      string     `json:"token,omitempty"` // Used if AuthMethod is 'token'
	AuthMethod AuthMethod `json:"auth_method"`
	RoleID     string     `json:"role_id,omitempty"`    // For AppRole
	SecretID   string     `json:"secret_id,omitempty"`  // For AppRole
	Namespace  string     `json:"namespace,omitempty"`  // For Vault Enterprise
	MountPath  string     `json:"mount_path,omitempty"` // Overrides default KVEnginePath
	CACert     string     `json:"ca_cert,omitempty"`    // Path to CA Cert for TLS
}

// KeyRotationPolicy defines rules for rotating encryption keys.
type KeyRotationPolicy struct {
	RotationInterval time.Duration `json:"rotation_interval"`
	Algorithm        string        `json:"algorithm"` // e.g., "AES-256-GCM"
	AutoRotate       bool          `json:"auto_rotate"`
}
