package main

import (
	"temp/config"
	route "temp/internal"
	"temp/internal/repository"
)

func main() {
	config.Run("config.yaml")
	repository.ConnectDatabase()
	engine := route.GenerateRouter()
	engine.Run(":8080")
	
}
