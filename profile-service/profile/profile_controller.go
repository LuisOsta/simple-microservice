package profile

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProfileBody struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
	UserId  string `json:"userId"`
}

type ProfileCreator = func(address string, phone string, uid string) (ProfileDocument, error)
type ProfileUpdator = func(uid string, p updatePayload) (ProfileDocument, error)
type Profile struct {
	CreateProfile ProfileCreator
	UpdateProfile ProfileUpdator
}

func (p *Profile) HandleCreateProfile(c *gin.Context) {

	var body createProfileBody
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	profile, err := p.CreateProfile(body.Address, body.Phone, body.UserId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

type updateProfileBody struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (p *Profile) HandleUpdateProfile(c *gin.Context) {
	uid := c.Param("userId")
	var body updateProfileBody
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	profile, err := p.UpdateProfile(uid, updatePayload(body))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, profile)
	}
}
