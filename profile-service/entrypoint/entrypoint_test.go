package entrypoint

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type signatureTest struct {
	host               string
	expectedStatusCode int
}

func TestCheckRequestSignature(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	tests := []signatureTest{
		{host: "localhost:5000", expectedStatusCode: http.StatusOK},
		{host: "wronghost:8000", expectedStatusCode: http.StatusUnauthorized},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/", io.Reader(strings.NewReader("")))
		c.Request.Host = test.host

		CheckRequestSignature(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d for host %s, got %d", test.expectedStatusCode, test.host, w.Code)
		}
	}
}
