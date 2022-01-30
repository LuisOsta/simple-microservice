package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type loginTest struct {
	expectedStatusCode   int
	credential_retriever CredentialRetriever
	body                 string
}

func mock_failed_get_user_by_credentials(username string, password string) (AuthDocument, error) {
	return AuthDocument{}, fmt.Errorf("Failed to get user by credentials")
}

func mock_successful_get_user_by_credentials(username string, password string) (AuthDocument, error) {
	return AuthDocument{}, nil
}

func TestHandleLogin(t *testing.T) {
	handleLoginTests := []loginTest{
		{
			body:                 `{"username":"test","password":"test"}`,
			expectedStatusCode:   200,
			credential_retriever: mock_successful_get_user_by_credentials,
		},
		{
			body:                 `{"username":"test","password":"test"}`,
			expectedStatusCode:   401,
			credential_retriever: mock_failed_get_user_by_credentials,
		},
		{
			body:                 `"username":"test","password":"test"`,
			expectedStatusCode:   400,
			credential_retriever: mock_failed_get_user_by_credentials,
		},
	}

	for _, test := range handleLoginTests {
		auth := Auth{
			GetUserByCredentials: test.credential_retriever,
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/auth/login", strings.NewReader(string(test.body)))
		c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

		auth.HandleLogin(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", test.expectedStatusCode, w.Code)
		}

	}
}
