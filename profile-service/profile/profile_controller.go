package profile

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleCreateProfile(c *gin.Context) {

	profile, err := createProfile("placeholder", "placeholder")

	if err != nil {
		log.Println(err)
		c.JSON(503, gin.H{
			"message": "NOT IMPLEMENTED",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    profile,
		})
	}
}

func HandleUpdateProfile(c *gin.Context) {
	c.JSON(503, gin.H{
		"message": "NOT IMPLEMENTED",
	})
}
