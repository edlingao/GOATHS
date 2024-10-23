package models

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ApiKey struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateAPIKey(username string) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	claims := ApiKey{
		username,
		jwt.RegisteredClaims{
			Subject: username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	signed_api, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return signed_api, nil
}

func VerifyAPIKey(apiKey string) bool {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(apiKey, &ApiKey{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

  subject, _ := token.Claims.GetSubject()

  log.Print("TOKEN SUBJECT", subject) 

	return token.Valid
}

func GetUsernameFromToken(apiKey string) ( string, error ) {
  mySigningKey := []byte(os.Getenv("JWT_SECRET"))

  token, err := jwt.ParseWithClaims(apiKey, &ApiKey{}, func(token *jwt.Token) (interface{}, error) {
    return mySigningKey, nil
  })

  if err != nil {
    return "", err
  }

  username, err := token.Claims.GetSubject()

  return username, nil
}

