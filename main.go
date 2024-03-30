package main

import (
	app "artha-api/src"
	"artha-api/src/configs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("DOCKER_ENVIRONMENT") == "true" {
		log.Println("Running inside a Docker container")
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	gin.SetMode(configs.AppConfig().GIN_MODE)
	server := gin.Default()

	app.DBConnection()
	app.Middlewares(server)
	app.Routes(server)
	app.Serve(server, configs.AppConfig().PORT)
}
