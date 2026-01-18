package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

//
// ============================================================
// JSON HELPERS
// ============================================================
//

// MarshalJSON marshals with optional HTML escaping disabled.
func MarshalJSON(v any, disableEscape bool) ([]byte, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(!disableEscape)

	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

func EncodeJSONB(v any) pqtype.NullRawMessage {
	if v == nil {
		return pqtype.NullRawMessage{}
	}

	b, err := json.Marshal(v)
	if err != nil {
		return pqtype.NullRawMessage{}
	}

	return pqtype.NullRawMessage{RawMessage: b, Valid: true}
}

func DecodeJSONB[T any](src pqtype.NullRawMessage) (T, error) {
	var out T
	if !src.Valid {
		return out, nil
	}
	err := json.Unmarshal(src.RawMessage, &out)
	return out, err
}

func JSONBToValue[T any](src pqtype.NullRawMessage) T {
	var out T
	if src.Valid {
		_ = json.Unmarshal(src.RawMessage, &out)
	}
	return out
}

//
// ============================================================
// SLICE HELPERS
// ============================================================
//

// Contains reports whether a slice contains an item
func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

//
// ============================================================
// NULL MAPPERS (Typed — safe & idiomatic)
// ============================================================
//

// ---- ToNull (Go → sql.Null*) ----

func ToNullString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: v != ""}
}

func ToNullBool(v bool) sql.NullBool {
	return sql.NullBool{Bool: v, Valid: true}
}

func ToNullInt(v int64) sql.NullInt64 {
	return sql.NullInt64{Int64: v, Valid: v != 0}
}

func ToNullInt32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: v != 0}
}

func ToNullFloat(v float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: v, Valid: v != 0}
}

func ToNullTime(v time.Time) sql.NullTime {
	if v.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: v, Valid: true}
}

func ToNullUUID(v uuid.UUID) uuid.NullUUID {
	if v == uuid.Nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{UUID: v, Valid: true}
}

// ---- NullToValue (sql → Go value) ----

func NullStringToValue(v sql.NullString) string {
	if !v.Valid {
		return ""
	}
	return v.String
}

func NullBoolToValue(v sql.NullBool) bool {
	return v.Valid && v.Bool
}

func NullIntToValue(v sql.NullInt64) int64 {
	if !v.Valid {
		return 0
	}
	return v.Int64
}

func NullInt32ToValue(v sql.NullInt32) int32 {
	if !v.Valid {
		return 0
	}
	return v.Int32
}

func TimeOrZero(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func NullTimeToValue(v sql.NullTime) time.Time {
	if !v.Valid {
		return time.Time{}
	}
	return v.Time
}

func NullUUIDToValue(v uuid.NullUUID) uuid.UUID {
	if !v.Valid {
		return uuid.Nil
	}
	return v.UUID
}

// ---- NullToPtr (sql → *Go) ----

func NullStringToPtr(v sql.NullString) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

func NullBoolToPtr(v sql.NullBool) *bool {
	if !v.Valid {
		return nil
	}
	return &v.Bool
}

func NullIntToPtr(v sql.NullInt64) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}

func NullFloatToPtr(v sql.NullFloat64) *float64 {
	if !v.Valid {
		return nil
	}
	return &v.Float64
}

func NullTimeToPtr(v sql.NullTime) *time.Time {
	if !v.Valid {
		return nil
	}
	return &v.Time
}

func NullUUIDToPtr(v uuid.NullUUID) *uuid.UUID {
	if !v.Valid {
		return nil
	}
	return &v.UUID
}

//
// ============================================================
// UUID HELPERS
// ============================================================
//

func UUIDPtr(u uuid.UUID) *uuid.UUID {
	if u == uuid.Nil {
		return nil
	}
	return &u
}

func UUIDFromPtr(p *uuid.UUID) uuid.UUID {
	if p == nil {
		return uuid.Nil
	}
	return *p
}

func UUIDToString(p *uuid.UUID) string {
	if p == nil {
		return ""
	}
	return p.String()
}

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

//
// ============================================================
// MISC HELPERS
// ============================================================
//

func StrOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func DerefTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func DerefUUID(u *uuid.UUID) uuid.UUID {
	if u == nil {
		return uuid.Nil
	}
	return *u
}
