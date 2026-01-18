package helper

import (
	"crypto/rand"
	"errors"
	"sync"
	"time"
)

// -------------------- OTP STORAGE -------------------- //

type OTPEntry struct {
	Code      string
	CreatedAt time.Time
	ExpiresAt time.Time
}

var (
	otpStore = make(map[string]OTPEntry)
	mutex    sync.RWMutex
)

// -------------------- OTP FUNCTIONS -------------------- //

// GenerateRandomOTP generates a secure numeric OTP of given length (default 6 digits)
func GenerateRandomOTP(length int) string {
	if length <= 0 {
		length = 6
	}

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		// fallback in case of error
		for i := range b {
			b[i] = byte(i % 10)
		}
	}

	otp := make([]byte, length)
	for i, by := range b {
		otp[i] = '0' + (by % 10)
	}
	return string(otp)
}

// StoreOTP stores an OTP for a username with TTL
func StoreOTP(username, otp string, ttl time.Duration) error {
	if username == "" || otp == "" {
		return errors.New("username and otp cannot be empty")
	}

	entry := OTPEntry{
		Code:      otp,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
	}

	mutex.Lock()
	otpStore[username] = entry
	mutex.Unlock()
	return nil
}

// GetStoredOTP retrieves OTP code and creation time for a username
// Returns error if not found or expired (automatically deletes expired OTP)
func GetStoredOTP(username string) (code string, createdAt time.Time, err error) {
	mutex.RLock()
	entry, ok := otpStore[username]
	mutex.RUnlock()

	if !ok {
		return "", time.Time{}, errors.New("otp not found")
	}

	if time.Now().After(entry.ExpiresAt) {
		// delete expired OTP
		DeleteOTP(username)
		return "", entry.CreatedAt, errors.New("otp expired")
	}

	return entry.Code, entry.CreatedAt, nil
}

// DeleteOTP removes an OTP for a username
func DeleteOTP(username string) {
	mutex.Lock()
	delete(otpStore, username)
	mutex.Unlock()
}

// -------------------- OPTIONAL CLEANUP -------------------- //

// CleanupExpiredOTPs can be run periodically to clean up expired OTPs
func CleanupExpiredOTPs() {
	now := time.Now()
	mutex.Lock()
	for username, entry := range otpStore {
		if now.After(entry.ExpiresAt) {
			delete(otpStore, username)
		}
	}
	mutex.Unlock()
}
