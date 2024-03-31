package authService

import (
	"artha-api/src/database"
	"artha-api/src/helpers"
	"artha-api/src/models"
	authRequest "artha-api/src/requests/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(context *gin.Context, request *authRequest.LoginRequest) *models.Users {
	var user models.Users

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	filter := bson.M{"email": request.Email, "password": request.Password}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse("Wrong email or password", http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.HttpResponse(err.Error(), http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == "unverified" {
		helpers.HttpResponse("The email has not verified yet.", http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}
