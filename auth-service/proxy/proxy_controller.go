package proxy

import (
	"log"

	"github.com/gin-gonic/gin"
)

// gets the prefix of the path, determines if the path is to a supported service
// If its supported, then it will return the service name and the path without the prefix
// It will then send a request using the Service Endpoint and the path
func HandleProxyRequest(c *gin.Context) {
	path := c.Param("path")
	serviceName, path := getServiceNameAndPath(path)

	service, err := getService(serviceName)
	if err != nil {
		log.Printf("Attempted to request path %s for invalid service %s\n", path, serviceName)
		c.JSON(404, gin.H{"error": "Service not found"})
		return
	}

	res, err := SendServiceRequest(service.Endpoint, path)

	if err != nil {
		log.Printf("Error sending request to service %s: %s\n", serviceName, err.Error())
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(500, gin.H{
		"message": res,
	})
}
