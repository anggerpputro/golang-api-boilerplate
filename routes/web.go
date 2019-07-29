package routes

import (
	"../app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(router *gin.Engine) *gin.Engine {
	helloController := controllers.NewHelloController()

	router.GET("/", helloController.Hello)
	router.GET("/ping", helloController.Ping)

	return router
}
