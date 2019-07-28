package routes

import (
	"../app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupApiRoutes() *gin.Engine {
	r := gin.Default()

	helloController := controllers.NewHelloController()

	r.GET("/", helloController.Hello)
	r.GET("/ping", helloController.Ping)

	return r
}
