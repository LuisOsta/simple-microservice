package auth

import (
	"github.com/auth-service/config"
	"github.com/golang-jwt/jwt/v4"
)

const AUTH_HEADER = "Authorization"

func createAuthToken(uid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uid,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetConfiguration().JWT_SECRET))

	return tokenString, err
}
