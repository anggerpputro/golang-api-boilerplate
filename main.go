package main

import (
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupWebRoutes(router)

	apiRoutes := router.Group("/api")
	{
		routes.SetupApiRoutes(apiRoutes)
	}

	router.Run(":8080")
}
