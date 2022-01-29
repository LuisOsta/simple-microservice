package proxy

import (
	"io"
	"log"
	"net/http"
)

func SendServiceRequest(serviceEndpoint string, path string, method string, body io.Reader) (*http.Response, error) {
	log.Printf("Sending %s request to service %s\n at path %s", method, serviceEndpoint, path)
	client := &http.Client{}
	req, err := http.NewRequest("GET", serviceEndpoint+path, body)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	return response, err
}
