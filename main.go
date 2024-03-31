package main

import (
	app "artha-api/src"
	"artha-api/src/configs"
	"artha-api/src/helpers"
	"fmt"
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

	claims, err := helpers.VerifyJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGFydGhhLmJoYWt0aWJ1YW5hLmNvbSIsImV4cCI6MTcxNDQ5NzM0MywiaWQiOiI2NjA4OGNlZWI1NzgzZTQzMTE5MTgyZGUiLCJuYW1lIjoiQXJ0aGEgQWRtaW4iLCJ0YWciOiIwMDAwMCIsInVzZXJuYW1lIjoiQXJ0aGEifQ.zDpk8qimDvFBplLiawKx8Jbdzl0gZgce7thYeoP15_0")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(err)
		fmt.Println(claims)
	}

	app.DBConnection(configs.DBConfig().DB_DSN, configs.DBConfig().DB_DATABASE)
	app.Middlewares(server)
	app.Routes(server)
	app.Serve(server, configs.AppConfig().PORT)
}
