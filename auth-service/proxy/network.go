package proxy

import (
	"io"
	"log"
	"net/http"
)

func SendServiceRequest(serviceEndpoint string, method string, body io.Reader) (*http.Response, error) {
	log.Printf("Sending %s request to endpoint %s\n", method, serviceEndpoint)
	client := &http.Client{}
	req, err := http.NewRequest(method, serviceEndpoint, body)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	return response, err
}
