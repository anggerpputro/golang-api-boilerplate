package routes

import (
	"github.com/cone-partij/golang-api-boilerplate/app/controllers/api"
	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	auth := router.Group("/auth")
	{
		authController := api.NewAuthController()

		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	users := router.Group("/users")
	{
		userController := api.NewUserController()

		users.GET("/", userController.Index)
		users.GET("/:id", userController.Show)

		users.POST("/", userController.Store)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Destroy)
	}

	return router
}
