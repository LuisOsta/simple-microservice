package proxy

import (
	"io"
	"log"
	"net/http"
)

// Generic proxy function that will send the request to the target service based on the method, endpoint and body given.
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
