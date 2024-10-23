package models

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func EncryptPassword(password string) string {
	plain_password := []byte(password)
	hash := sha256.Sum256(plain_password)

	return hex.EncodeToString(hash[:])
}

func ValidatePassword(hash string, password string) bool {
	return EncryptPassword(password) == hash
}

func GetCookie(signed_api string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    signed_api,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	return cookie
}

func RemoveCookie() *http.Cookie {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	return cookie
}
