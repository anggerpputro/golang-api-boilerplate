package routes

import (
	"../app/controllers/api"
	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	auth := router.Group("/auth")
	{
		authController := apicontrollers.NewAuthController()

		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	return router
}
