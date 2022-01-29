package profile

import "github.com/gin-gonic/gin"

func HandleProfileRequest(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "NOT IMPLEMENTED",
	})
}
