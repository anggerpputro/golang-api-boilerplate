package main

import (
	"fmt"

	"github.com/cone-partij/golang-api-boilerplate/routes"
	"github.com/cone-partij/golang-api-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupWebRoutes(router)

	apiRoutes := router.Group("/api")
	{
		routes.SetupApiRoutes(apiRoutes)
	}

	fmt.Printf("\nENV:\n")
	fmt.Printf("- DB_CONNECTION \t: %s \n", utils.Env("DB_CONNECTION"))
	fmt.Printf("- DB_HOST \t\t: %s \n", utils.Env("DB_HOST"))
	fmt.Printf("- DB_DATABASE \t\t: %s \n", utils.Env("DB_DATABASE"))
	fmt.Printf("- DB_USERNAME \t\t: %s \n", utils.Env("DB_USERNAME"))
	fmt.Printf("- DB_PASSWORD \t\t: %s \n", utils.Env("DB_PASSWORD"))
	fmt.Printf("\n")

	router.Run(":8080")
}
