package profile

import (
	"log"

	"github.com/gin-gonic/gin"
)

type createProfileBody struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
	UserId  string `json:"userId"`
}

func HandleCreateProfile(c *gin.Context) {

	var body createProfileBody
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	profile, err := createProfile(body.Address, body.Phone, body.UserId)

	if err != nil {
		log.Println(err)
		c.JSON(503, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, profile)
	}
}

type updateProfileBody struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func HandleUpdateProfile(c *gin.Context) {
	uid := c.Param("userId")
	var body updateProfileBody
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	profile, err := updateProfile(uid, updatePayload(body))

	if err != nil {
		log.Println(err)
		c.JSON(503, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, profile)
	}
}
