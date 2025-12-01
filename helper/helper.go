package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func JSONMarshal(v map[string]string, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckString(statuscreatedat sql.NullString) string {
	statuscreatedat_ := ""
	if statuscreatedat.Valid {
		statuscreatedat_ = statuscreatedat.String
	}
	return statuscreatedat_
}

func UUIDToString(id *uuid.UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func ToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ToNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{
		UUID:  *u,
		Valid: true,
	}
}

// ToNullTime converts *time.Time or time.Time to sql.NullTime
func ToNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{Valid: false}
}

// ToNullBool converts bool to sql.NullBool
func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}
