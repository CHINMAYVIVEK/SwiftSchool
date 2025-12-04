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

// ----------------------------------------
// JSON / JSONB Helpers
// ----------------------------------------

// JSONMarshal safely marshals any value to JSON with optional unescaped characters
func JSONMarshal(v any, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	if safeEncoding {
		replacements := []struct {
			old []byte
			new []byte
		}{
			{[]byte("\\u003c"), []byte("<")},
			{[]byte("\\u003e"), []byte(">")},
			{[]byte("\\u0026"), []byte("&")},
		}
		for _, r := range replacements {
			b = bytes.ReplaceAll(b, r.old, r.new)
		}
	}
	return b, nil
}

// EncodeJSONB encodes any value into pqtype.NullRawMessage
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

// DecodeJSONB decodes pqtype.NullRawMessage into T (struct or slice)
func DecodeJSONB[T any](src pqtype.NullRawMessage) (T, error) {
	var out T
	if !src.Valid {
		return out, nil
	}
	err := json.Unmarshal(src.RawMessage, &out)
	return out, err
}

// JSONBToValue decodes pqtype.NullRawMessage into type T.
// Returns zero value of T if NullRawMessage is invalid.
func JSONBToValue[T any](src pqtype.NullRawMessage) T {
	var out T
	if !src.Valid {
		return out
	}
	_ = json.Unmarshal(src.RawMessage, &out) // optionally log error
	return out
}

// ----------------------------------------
// Slice Helpers
// ----------------------------------------

// Contains checks if a slice contains a given item
func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// ----------------------------------------
// NullString Helpers
// ----------------------------------------

// ToNullString converts string, string alias, or *string to sql.NullString
func ToNullString[T ~string | *string](s T) sql.NullString {
	switch v := any(s).(type) {
	case string:
		if v == "" {
			return sql.NullString{}
		}
		return sql.NullString{String: v, Valid: true}
	case *string:
		if v == nil || *v == "" {
			return sql.NullString{}
		}
		return sql.NullString{String: *v, Valid: true}
	default:
		return sql.NullString{}
	}
}

// NullToValue converts sql.Null* types to their Go value with default zero
func NullToValue[T any](v any) T {
	var zero T
	switch val := v.(type) {
	case sql.NullString:
		if !val.Valid {
			return zero
		}
		return any(val.String).(T)
	case sql.NullBool:
		if !val.Valid {
			return zero
		}
		return any(val.Bool).(T)
	case sql.NullInt64:
		if !val.Valid {
			return zero
		}
		return any(val.Int64).(T)
	case sql.NullFloat64:
		if !val.Valid {
			return zero
		}
		return any(val.Float64).(T)
	case sql.NullTime:
		if !val.Valid {
			return zero
		}
		return any(val.Time).(T)
	case uuid.NullUUID:
		if !val.Valid {
			return zero
		}
		return any(val.UUID).(T)
	default:
		return zero
	}
}

// NullToPointer converts sql.Null* types to *T safely, using the Valid field
func NullToPointer[T any](v any) *T {
	switch val := v.(type) {
	case sql.NullString:
		if !val.Valid {
			return nil
		}
		v := any(val.String).(T)
		return &v
	case sql.NullBool:
		if !val.Valid {
			return nil
		}
		v := any(val.Bool).(T)
		return &v
	case sql.NullInt64:
		if !val.Valid {
			return nil
		}
		v := any(val.Int64).(T)
		return &v
	case sql.NullFloat64:
		if !val.Valid {
			return nil
		}
		v := any(val.Float64).(T)
		return &v
	case sql.NullTime:
		if !val.Valid {
			return nil
		}
		v := any(val.Time).(T)
		return &v
	case uuid.NullUUID:
		if !val.Valid {
			return nil
		}
		v := any(val.UUID).(T)
		return &v
	default:
		return nil
	}
}

// ----------------------------------------
// UUID Helpers
// ----------------------------------------

// ToNullUUID converts *uuid.UUID or uuid.UUID to uuid.NullUUID
func ToNullUUID[T *uuid.UUID | uuid.UUID](id T) uuid.NullUUID {
	switch v := any(id).(type) {
	case *uuid.UUID:
		if v == nil {
			return uuid.NullUUID{}
		}
		return uuid.NullUUID{UUID: *v, Valid: true}
	case uuid.UUID:
		if v == uuid.Nil {
			return uuid.NullUUID{}
		}
		return uuid.NullUUID{UUID: v, Valid: true}
	default:
		return uuid.NullUUID{}
	}
}

// ToUUIDPointer converts uuid.UUID or uuid.NullUUID to *uuid.UUID
func ToUUIDPointer[T uuid.UUID | uuid.NullUUID](id T) *uuid.UUID {
	switch v := any(id).(type) {
	case uuid.UUID:
		return &v
	case uuid.NullUUID:
		if !v.Valid {
			return nil
		}
		return &v.UUID
	default:
		return nil
	}
}

// UUIDToString safely converts *uuid.UUID to string
func UUIDToString(id *uuid.UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

// StringToUUID parses string to uuid.UUID
func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// ----------------------------------------
// NullTime Helpers
// ----------------------------------------

// ToNullTime converts *time.Time to sql.NullTime
func ToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// ----------------------------------------
// NullBool Helpers
// ----------------------------------------

// ToNullBool converts bool or *bool to sql.NullBool
func ToNullBool[T bool | *bool](b T) sql.NullBool {
	switch v := any(b).(type) {
	case bool:
		return sql.NullBool{Bool: v, Valid: true}
	case *bool:
		if v == nil {
			return sql.NullBool{}
		}
		return sql.NullBool{Bool: *v, Valid: true}
	default:
		return sql.NullBool{}
	}
}

// ----------------------------------------
// Misc Helpers
// ----------------------------------------

// ToStr safely dereferences a *string or returns ""
func ToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
