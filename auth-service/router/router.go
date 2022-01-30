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

// The createProxyGroup function has to use 'NoRoute' instead of 'Any' to register the Proxy handler of the non-auth routes due to limitations of the Gin router.
// Its not able to handle complete wildcard routes and the auth routes at the same time.
// While conceptually awkward and unideal, due to the nature of the API gateway, we don't pay any flexibility or performance cost for this limitation.
func createProxyGroup(router *gin.Engine) {
	p := proxy.Proxy{SendServiceRequest: proxy.SendServiceRequest}
	router.Use(auth.CheckAuthentication)
	router.NoRoute(p.HandleProxyRequest)
}

func createAuthGroup(router *gin.Engine) {
	authGroup := router.Group("/auth")
	a := auth.Auth{GetUserByCredentials: auth.GetUserByCredentials}
	authGroup.POST("/login", a.HandleLogin)
}
