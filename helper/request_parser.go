package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// ParseUUIDFromQuery parses a UUID from the query parameters
func ParseUUIDFromQuery(r *http.Request, key string) (uuid.UUID, error) {
	idStr := r.URL.Query().Get(key)
	if idStr == "" {
		return uuid.Nil, nil
	}
	return uuid.Parse(idStr)
}

// ParseRequiredUUIDFromQuery parses a required UUID from the query parameters
func ParseRequiredUUIDFromQuery(r *http.Request, key string) (uuid.UUID, error) {
	idStr := r.URL.Query().Get(key)
	if idStr == "" {
		return uuid.Nil, ErrMissingParameter(key)
	}
	return uuid.Parse(idStr)
}

// ParseUUIDFromPath parses a UUID from the URL path variables
func ParseUUIDFromPath(r *http.Request, key string) (uuid.UUID, error) {
	vars := mux.Vars(r)
	idStr, ok := vars[key]
	if !ok || idStr == "" {
		return uuid.Nil, ErrMissingParameter(key)
	}
	return uuid.Parse(idStr)
}

// DecodeJSONBody decodes the JSON body into the target struct
func DecodeJSONBody(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return ErrEmptyRequestBody
	}
	return json.NewDecoder(r.Body).Decode(target)
}

// GetRequiredQueryParam returns a required query parameter
func GetRequiredQueryParam(r *http.Request, key string) (string, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return "", ErrMissingParameter(key)
	}
	return val, nil
}

// GetQueryParam returns an optional query parameter with a default value
func GetQueryParam(r *http.Request, key, defaultValue string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// GetInstituteID extracts the Institute ID from the request context or headers
// Priority:
// 1. Context (set by middleware) - TODO
// 2. Header (X-Institute-ID) - For testing/API keys
func GetInstituteID(r *http.Request) (uuid.UUID, error) {
	// 1. Try Header
	idStr := r.Header.Get("X-Institute-ID")
	if idStr != "" {
		return uuid.Parse(idStr)
	}

	// 2. Try Context (Placeholder for when we implement session middleware fully)
	// val := r.Context().Value("institute_id")
	// if id, ok := val.(uuid.UUID); ok {
	// 	return id, nil
	// }

	return uuid.Nil, errors.New("institute ID not found in request (missing X-Institute-ID header)")
}
