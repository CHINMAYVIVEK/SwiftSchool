package helper

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                    TYPES                        //
//////////////////////////////////////////////////////

type SessionData struct {
	UserID   string
	Username string
	Role     string
	Expiry   time.Time
}

//////////////////////////////////////////////////////
//                SESSION STORE                   //
//////////////////////////////////////////////////////

type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*SessionData
}

var store = &SessionStore{
	sessions: make(map[string]*SessionData),
}

const (
	SessionCookieName = "swiftschool_session"
	SessionDuration   = 24 * time.Hour
)

//////////////////////////////////////////////////////
//                 CONTEXT KEY                    //
//////////////////////////////////////////////////////

type sessionContextKeyType struct{}

var sessionContextKey = sessionContextKeyType{}

//////////////////////////////////////////////////////
//              SESSION OPERATIONS                //
//////////////////////////////////////////////////////

func CreateSession(
	w http.ResponseWriter,
	userID, username, role string,
) error {
	sessionID := uuid.NewString()

	session := &SessionData{
		UserID:   userID,
		Username: username,
		Role:     role,
		Expiry:   time.Now().Add(SessionDuration),
	}

	store.mu.Lock()
	store.sessions[sessionID] = session
	store.mu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.Expiry,
	})

	return nil
}

func GetSession(r *http.Request) (*SessionData, error) {
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return nil, errors.New("no session cookie")
	}

	store.mu.RLock()
	session, ok := store.sessions[cookie.Value]
	store.mu.RUnlock()

	if !ok {
		return nil, errors.New("invalid session")
	}

	if time.Now().After(session.Expiry) {
		DeleteSession(cookie.Value)
		return nil, errors.New("session expired")
	}

	return session, nil
}

func DeleteSession(sessionID string) {
	store.mu.Lock()
	delete(store.sessions, sessionID)
	store.mu.Unlock()
}

//////////////////////////////////////////////////////
//                CONTEXT HELPERS                 //
//////////////////////////////////////////////////////

func WithSession(ctx context.Context, session *SessionData) context.Context {
	return context.WithValue(ctx, sessionContextKey, session)
}

func SessionFromContext(ctx context.Context) (*SessionData, bool) {
	session, ok := ctx.Value(sessionContextKey).(*SessionData)
	return session, ok
}

//////////////////////////////////////////////////////
//                 MIDDLEWARE                    //
//////////////////////////////////////////////////////

func RequireSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSession(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := WithSession(r.Context(), session)
		next(w, r.WithContext(ctx))
	}
}

//////////////////////////////////////////////////////
//               DEBUG / UTILS                    //
//////////////////////////////////////////////////////

func SessionToJSON(session *SessionData) string {
	b, _ := json.Marshal(session)
	return string(b)
}
