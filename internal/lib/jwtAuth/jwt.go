package jwtauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateJWTToken generates a JWT token for the given user ID
func GenerateJWTToken(userID uint64, email, secret string, expires time.Duration) (string, error) {
	expirationTime := time.Now().Add(expires).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     expirationTime,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWTToken verifies and parses the JWT token
func VerifyJWTToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// GetUserIDFromJWT extracts the user ID from a JWT token.
func GetUserIDFromJWT(tokenString string, secret string) (uint64, error) {
	token, err := VerifyJWTToken(tokenString, secret)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID in token claims")
	}

	return uint64(userID), nil
}
