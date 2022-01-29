package router

import (
	"github.com/user-service/profile"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/", profile.HandleCreateProfile)

	router.PUT("/:userId", profile.HandleUpdateProfile)

	return router
}
