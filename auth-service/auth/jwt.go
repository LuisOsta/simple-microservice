package auth

import (
	"strings"

	"github.com/auth-service/config"
	"github.com/gin-gonic/gin"
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

func isTokenValid(tokenString string) bool {

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfiguration().JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	return err == nil
}

// Checks the AUTH_HEADER for a valid token, parses the Bearer token from the value of the header. Then verifies whether or not its valid.
func CheckAuthentication(c *gin.Context) {

	if c.GetHeader(AUTH_HEADER) == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	rawToken := strings.Split(c.GetHeader(AUTH_HEADER), " ")[1]
	hasPermission := isTokenValid(rawToken)

	if !hasPermission {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
	} else {
		c.Next()
	}
}
