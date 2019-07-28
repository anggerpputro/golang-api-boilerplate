package main

import "./routes"

func main() {
	engine := routes.SetupApiRoutes()

	engine.Run(":8080")
}
