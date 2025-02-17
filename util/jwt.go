package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var hmacKey = "secret"

func GenerateJWT(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(hmacKey))
	return tokenString
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unsupported signing method")
		}
		return []byte(hmacKey), nil
	})
	if err != nil {
		return "", err
	}
	userID := token.Claims.(jwt.MapClaims)["sub"].(string)
	return userID, nil
}
