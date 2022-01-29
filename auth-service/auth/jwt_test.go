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

	var checkAuthTests = []CheckAuthTest{
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
