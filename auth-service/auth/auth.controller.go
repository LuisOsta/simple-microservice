package auth

import "github.com/gin-gonic/gin"

func HandleLogin(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "NOT IMPLEMENTED",
	})
}
