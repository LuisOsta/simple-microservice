package profile

import (
	"log"

	"github.com/gin-gonic/gin"
)

type createProfileBody struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func HandleCreateProfile(c *gin.Context) {

	var body createProfileBody
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	profile, err := createProfile(body.Address, body.Phone)

	if err != nil {
		log.Println(err)
		c.JSON(503, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, profile)
	}
}

func HandleUpdateProfile(c *gin.Context) {
	c.JSON(503, gin.H{
		"message": "NOT IMPLEMENTED",
	})
}
