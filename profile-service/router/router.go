package router

import (
	"github.com/user-service/entrypoint"
	"github.com/user-service/profile"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(entrypoint.CheckRequestSignature)

	p := profile.Profile{CreateProfile: profile.CreateProfile, UpdateProfile: profile.UpdateProfile}

	router.POST("/", p.HandleCreateProfile)

	router.PUT("/:userId", p.HandleUpdateProfile)

	return router
}
