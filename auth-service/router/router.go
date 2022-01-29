package router

import (
	"github.com/auth-service/auth"
	"github.com/auth-service/proxy"
	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	var groupConfigurers []func(*gin.Engine)
	groupConfigurers = append(groupConfigurers, createAuthGroup)
	groupConfigurers = append(groupConfigurers, createProxyGroup)

	for _, fn := range groupConfigurers {
		fn(router)
	}

	return router
}

func createProxyGroup(router *gin.Engine) {
	profileGroup := router.Group("/")
	profileGroup.Use(auth.CheckAuthentication)
	profileGroup.Any("/:path", proxy.HandleProxyRequest)
}

func createAuthGroup(router *gin.Engine) {
	authGroup := router.Group("/auth")

	authGroup.POST("/login", auth.HandleLogin)
}
