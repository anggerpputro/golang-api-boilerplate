package main

import (
	"fmt"

	ctlModel "github.com/cone-partij/golang-api-boilerplate/app/models"
	ctlDB "github.com/cone-partij/golang-api-boilerplate/database"
	ctlRouter "github.com/cone-partij/golang-api-boilerplate/routes"
	ctlUtil "github.com/cone-partij/golang-api-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ctlRouter.SetupWebRoutes(router)

	apiRoutes := router.Group("/api")
	{
		ctlRouter.SetupApiRoutes(apiRoutes)
	}

	fmt.Printf("\nENV:\n")
	fmt.Printf("- DB_CONNECTION \t: %s \n", ctlUtil.Env("DB_CONNECTION"))
	fmt.Printf("- DB_HOST \t\t: %s \n", ctlUtil.Env("DB_HOST"))
	fmt.Printf("- DB_DATABASE \t\t: %s \n", ctlUtil.Env("DB_DATABASE"))
	fmt.Printf("- DB_USERNAME \t\t: %s \n", ctlUtil.Env("DB_USERNAME"))
	fmt.Printf("- DB_PASSWORD \t\t: %s \n", ctlUtil.Env("DB_PASSWORD"))
	fmt.Printf("\n")

	fmt.Printf("\nConnecting to Database...\n")

	db := ctlDB.GetDefaultConnection()
	defer db.Close()

	fmt.Printf("Connected to Database!\n\n")

	ctlModel.Migrate(db)

	router.Run(":8080")
}
