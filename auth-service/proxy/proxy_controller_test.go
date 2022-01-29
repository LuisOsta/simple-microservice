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
}

func TestHandleProxyRequest(t *testing.T) {

	handleProxyTests := []HandleProxyTest{
		{
			path:                   "/profile",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_successful_send_service_request,
		},
		{
			path:                   "/profile",
			method:                 "POST",
			body:                   "{}",
			service_request_sender: mock_failed_send_service_request,
		},
		{
			path:                   "/profile/12345",
			method:                 "PUT",
			body:                   "{}",
			service_request_sender: mock_successful_send_service_request,
		},
		{
			path:                   "/profile/12345",
			method:                 "PUT",
			body:                   "{}",
			service_request_sender: mock_failed_send_service_request,
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
	}

}

func BenchmarkHandleProxyRequest(t *testing.B) {

}
