package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/auth-service/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const AUTH_HEADER = "Authorization"

// Uses the HS256 algorithm to generate a JWT token based on the userId and the JWT secret.
func createAuthToken(uid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uid,
	})

	// Sign and get the complete encoded token as a string using the secret.
	tokenString, err := token.SignedString([]byte(config.GetConfiguration().JWT_SECRET))

	return tokenString, err
}

// The token should only be recognized as valid if it was generated with the correct algorithm and secret.
func isTokenValid(tokenString string) bool {

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfiguration().JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	return err == nil
}

// Checks the AUTH_HEADER for a valid token, parses the Bearer token from the value of the header. Then verifies whether or not its valid.
// We want to make sure that only Bearer tokens generated with the appropriate algorithm and secret are allowed to continue.
func CheckAuthentication(c *gin.Context) {
	tokenString := c.GetHeader(AUTH_HEADER)
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		log.Printf("No token found in request\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenParts := strings.Split(tokenString, " ")

	if tokenParts[0] != "Bearer" {
		log.Printf("Invalid token format\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	} else if !isTokenValid(tokenParts[1]) {
		log.Printf("Invalid token: %s\n", tokenParts[1])
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	} else {
		c.Next()
	}
}
