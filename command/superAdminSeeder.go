package main

import (
	"artha-api/src/configs"
	"artha-api/src/database"
	"artha-api/src/helpers"
	"artha-api/src/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(configs.DBConfig().DB_DSN, configs.DBConfig().DB_DATABASE); err != nil {
		fmt.Println(err)
		return
	}

	superAdminUser := models.Users{
		Name:        "Artha Admin",
		Username:    "Artha",
		Tag:         "00000",
		Email:       "admin@artha.bhaktibuana.com",
		Password:    helpers.HashPassword(os.Getenv("SUPER_ADMIN_PASSWORD")),
		Birthdate:   time.Date(1998, 12, 9, 0, 0, 0, 0, time.UTC),
		Gender:      models.USER_GENDER_MALE,
		AccountType: models.USER_ACCOUNT_TYPE_ADMIN,
		Status:      models.USER_STATUS_VERIFIED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := database.Users.InsertOne(context.Background(), superAdminUser); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Super admin user seeded successfully!")

	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
