package router

import (
	"github.com/user-service/controller"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/", controller.HandleCreateCustomer)

	router.PUT("/:userId", controller.HandleUpdateCustomer)

	return router
}
