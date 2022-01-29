package proxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func mock_failed_send_service_request(serviceEndpoint string, method string, body io.Reader) (*http.Response, error) {
	return nil, fmt.Errorf("Failed to send request to service")
}

func mock_successful_send_service_request(serviceEndpoint string, method string, body io.Reader) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("{}")),
	}, nil
}

type HandleProxyTest struct {
	path                   string
	method                 string
	body                   string
	service_request_sender ServiceRequestSender
	expectedStatusCode     int
}

func TestHandleProxyRequest(t *testing.T) {

	handleProxyTests := []HandleProxyTest{
		{
			path:                   "/profile",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_successful_send_service_request,
			expectedStatusCode:     200,
		},
		{
			path:                   "/profile",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_failed_send_service_request,
			expectedStatusCode:     500,
		},
		{
			path:                   "/profile/12345",
			method:                 "PUT",
			body:                   "{}",
			service_request_sender: mock_successful_send_service_request,
			expectedStatusCode:     200,
		},
		{
			path:                   "/profile/12345",
			method:                 "PUT",
			body:                   "{}",
			service_request_sender: mock_failed_send_service_request,
			expectedStatusCode:     500,
		},
		{
			path:                   "/location",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_successful_send_service_request,
			expectedStatusCode:     404,
		},
		{
			path:                   "/location",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_failed_send_service_request,
			expectedStatusCode:     404,
		},
	}

	for _, test := range handleProxyTests {
		proxy := Proxy{
			SendServiceRequest: test.service_request_sender,
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(test.method, test.path, io.Reader(strings.NewReader(test.body)))
		c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

		proxy.HandleProxyRequest(c)

		if w.Code != test.expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", test.expectedStatusCode, w.Code)
		}
	}

}

func mock_benched_send_service_request(serviceEndpoint string, method string, body io.Reader) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("{}")),
	}, nil
}

func BenchmarkHandleProxyRequest(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	for i := 0; i < b.N; i++ {
		proxy := Proxy{
			SendServiceRequest: mock_benched_send_service_request,
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/profile", io.Reader(strings.NewReader("{}")))
		c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

		proxy.HandleProxyRequest(c)

		if w.Code != 200 {
			b.Errorf("Expected status code %d, got %d", 200, w.Code)
		}
	}
}
