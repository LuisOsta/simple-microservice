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

	router.POST("/", profile.HandleCreateProfile)

	router.PUT("/:userId", profile.HandleUpdateProfile)

	return router
}
