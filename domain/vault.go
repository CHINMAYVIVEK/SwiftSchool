package domain

// -------------------------
// Vault / KV Configuration
// -------------------------
import "fmt"

var (
	KVEnginePath       = "secret"         // Default Vault KV engine path
	InstituteSecretKey = "institute-keys" // Secret path holding all institute keys
	FieldKey           = "key"            // field prefix inside institute-keys
	DefaultAESKeyLen   = 32               // Default AES key length in bytes
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
// Errors
// -------------------------
var (
	ErrInvalidKeyLength = fmt.Errorf("encryption key must be %d bytes", DefaultAESKeyLen)
)
