package router

import (
	"github.com/auth-service/auth"
	"github.com/auth-service/profile"
	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	var groupConfigurers []func(*gin.Engine)
	groupConfigurers = append(groupConfigurers, createProfileGroup)
	groupConfigurers = append(groupConfigurers, createAuthGroup)

	for _, fn := range groupConfigurers {
		fn(router)
	}

	return router
}

func createProfileGroup(router *gin.Engine) {
	profileGroup := router.Group("/user")
	profileGroup.Use(auth.CheckAuthentication)
	profileGroup.Any("/:path", profile.HandleProfileRequest)
}

func createAuthGroup(router *gin.Engine) {
	authGroup := router.Group("/auth")

	authGroup.POST("/login", auth.HandleLogin)
}
