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

// ----------------------
// JSON / JSONB Helpers
// ----------------------

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

// DecodeJSONB decodes a pqtype.NullRawMessage into a slice of T
func DecodeJSONB[T any](src pqtype.NullRawMessage) ([]T, error) {
	if !src.Valid {
		return nil, nil
	}
	var out []T
	return out, json.Unmarshal(src.RawMessage, &out)
}

// EncodeJSONB encodes any value into pqtype.NullRawMessage
func EncodeJSONB(v any) pqtype.NullRawMessage {
	if v == nil {
		return pqtype.NullRawMessage{}
	}
	b, err := json.Marshal(v)
	if err != nil {
		return pqtype.NullRawMessage{} // optionally log the error
	}
	return pqtype.NullRawMessage{
		RawMessage: b,
		Valid:      true,
	}
}

// ----------------------
// Slice / Utility Helpers
// ----------------------

// Contains checks if a slice of any comparable type contains the given item
func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// ----------------------
// String / NullString Helpers
// ----------------------

// CheckString returns string value of sql.NullString or empty string
func CheckString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// NullStringToPtr converts sql.NullString to *string
func NullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// PtrToNullString converts *string to sql.NullString
func PtrToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// ----------------------
// UUID / NullUUID Helpers
// ----------------------

// PtrToNullUUID converts *uuid.UUID to uuid.NullUUID
func PtrToNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{UUID: *u, Valid: true}
}

// NullUUIDToPtr converts uuid.NullUUID to *uuid.UUID
func NullUUIDToPtr(u uuid.NullUUID) *uuid.UUID {
	if !u.Valid {
		return nil
	}
	return &u.UUID
}

// UUIDToPtr converts a UUID to a pointer
func UUIDToPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

// UUIDToString converts *uuid.UUID to string safely
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

// ----------------------
// Time / NullTime Helpers
// ----------------------

// PtrToNullTime converts *time.Time to sql.NullTime
func PtrToNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{Valid: false}
}

// NullTimeToPtr converts sql.NullTime to *time.Time
func NullTimeToPtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// ----------------------
// Bool / NullBool Helpers
// ----------------------

// PtrToNullBool converts *bool to sql.NullBool
func PtrToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

// NullBoolToPtr converts sql.NullBool to *bool
func NullBoolToPtr(nb sql.NullBool) *bool {
	if !nb.Valid {
		return nil
	}
	return &nb.Bool
}

// NullBoolToBool converts sql.NullBool to bool with default false
func NullBoolToBool(nb sql.NullBool) bool {
	return nb.Valid && nb.Bool
}

// BoolToNullBool converts bool to sql.NullBool (always valid)
func BoolToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

// ----------------------
// General Helpers
// ----------------------

// ToStr safely dereferences a string pointer or returns empty string
func ToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
