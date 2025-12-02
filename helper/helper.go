package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"slices"
	"time"

	"github.com/google/uuid"
)

// JSONMarshal safely marshals a map to JSON with optional unescaped characters
func JSONMarshal(v map[string]string, safeEncoding bool) ([]byte, error) {
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

// Contains checks if a slice of any comparable type contains the given item
func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// CheckString returns the value of sql.NullString or empty string if null
func CheckString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// UUIDToString converts a *uuid.UUID to string safely
func UUIDToString(id *uuid.UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

// StringToUUID converts a string to uuid.UUID
func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// ToStr safely dereferences a string pointer or returns empty string
func ToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ToNullUUID converts a *uuid.UUID to uuid.NullUUID
func ToNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{UUID: *u, Valid: true}
}

// ToNullTime converts a *time.Time to sql.NullTime
func ToNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{Valid: false}
}

// ToNullBool converts a bool to sql.NullBool
func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}
