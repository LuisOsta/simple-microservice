package proxy

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// gets the prefix of the path, determines if the path is to a supported service
// If its supported, then it will return the service name and the path without the prefix
// It will then send a request using the Service Endpoint and the path
func HandleProxyRequest(c *gin.Context) {
	serviceName, servicePath := getServiceNameAndPath(c.Request.URL.Path)
	service, err := getService(serviceName)
	if err != nil {
		log.Printf("Attempted to request path %s for invalid service %s\n", servicePath, serviceName)
		c.JSON(404, gin.H{"error": "Service not found"})
		return
	}

	res, err := SendServiceRequest(service.Endpoint+servicePath, c.Request.Method, c.Request.Body)

	if err != nil {
		log.Printf("Error sending request to service %s: %s\n", serviceName, err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("Error reading response from service %s: %s\n", serviceName, err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.Data(res.StatusCode, res.Header.Get("Content-Type"), body)
}
