package entrypoint

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user-service/config"
)

// A Gin middleware that verifies the request signature to confirm it came from the router
// In production and advanced use case, more sophisticated verification is needed.
func CheckRequestSignature(c *gin.Context) {
	if c.Request.Host != config.GetConfiguration().ROUTER_HOST {
		log.Println("Request originated from unauthorized host " + c.Request.Host)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid request"})
		c.Abort()
	} else {
		c.Next()
	}
}
