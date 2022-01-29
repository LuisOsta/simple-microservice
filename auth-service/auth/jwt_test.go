package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/auth-service/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CheckAuthTest struct {
	token              string
	expectedStatusCode int
}

func TestCheckAuthentication(t *testing.T) {
	uid := "1234"
	token, _ := createAuthToken(uid)
	otherToken, _ := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"userId": uid,
	}).SignedString([]byte(config.GetConfiguration().JWT_SECRET))
	unauthorizedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uid,
	}).SignedString([]byte("wrong secret"))

	checkAuthTests := []CheckAuthTest{
		{token: "", expectedStatusCode: http.StatusUnauthorized},
		{token: " ", expectedStatusCode: http.StatusUnauthorized},
		{token: token, expectedStatusCode: http.StatusUnauthorized},
		{token: otherToken, expectedStatusCode: http.StatusUnauthorized},
		{token: "Bearer " + unauthorizedToken, expectedStatusCode: http.StatusUnauthorized},
		{token: "Bearer " + token, expectedStatusCode: http.StatusOK},
	}
	for _, test := range checkAuthTests {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.Header.Set(AUTH_HEADER, test.token)
		CheckAuthentication(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d for token %s, got %d", test.expectedStatusCode, test.token, w.Code)
		}

	}

}

type TokenValidTest struct {
	token    string
	expected bool
}

func TestTokenValid(t *testing.T) {
	uid := "1234"
	token, _ := createAuthToken(uid)
	otherToken, _ := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"userId": uid,
	}).SignedString([]byte(config.GetConfiguration().JWT_SECRET))
	unauthorizedToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uid,
	}).SignedString([]byte("wrong secret"))

	tokenValidTests := []TokenValidTest{
		{token: otherToken, expected: false},
		{token: unauthorizedToken, expected: false},
		{token: token, expected: true},
	}

	for _, test := range tokenValidTests {
		result := isTokenValid(test.token)
		if result != test.expected {
			t.Errorf("Expected token %s to be %t, got %t", test.token, test.expected, result)
		}
	}
}

func TestCreateToken(t *testing.T) {

	uid := "1234"
	tokenString, _ := createAuthToken(uid)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfiguration().JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if uid != token.Claims.(jwt.MapClaims)["userId"] {
		t.Errorf("Expected token to contain userId %s, got %s", uid, token.Claims.(jwt.MapClaims)["userId"])
	}
}
