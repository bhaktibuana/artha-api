package authService

import (
	"artha-api/src/apis"
	"artha-api/src/configs"
	"artha-api/src/constants"
	"artha-api/src/database"
	"artha-api/src/helpers"
	"artha-api/src/models"
	authRequest "artha-api/src/requests/auth"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_LoginRequest
 * @returns *models.Users
 */
func Login(context *gin.Context, request *authRequest.S_LoginRequest) *models.Users {
	var user models.Users

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	filter := bson.M{"email": request.Email, "password": request.Password}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.WRONG_MAIL_PASS, http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == "unverified" {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// Register Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_RegisterRequest
 * @returns *models.Users
 */
func Register(context *gin.Context, request *authRequest.S_RegisterRequest) *models.Users {
	var user models.Users
	var username, tag, accountId string

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	for {
		username, tag, accountId = helpers.GenerateAccountId()

		filter := bson.M{"username": username, "tag": tag}

		if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
			if err == mongo.ErrNoDocuments {
				break
			}
		}
	}

	user = models.Users{
		ID:          primitive.NewObjectID(),
		Name:        request.Name,
		Username:    username,
		Tag:         tag,
		Email:       request.Email,
		Password:    request.Password,
		AccountType: "user",
		Status:      "unverified",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := database.Users.InsertOne(context, &user); err != nil {
		helpers.HttpResponse(constants.REGISTER_FAILED, http.StatusBadRequest, context, nil)
		return nil
	}

	claims := jwt.MapClaims{
		"id": user.ID,
	}

	token, _ := helpers.GenerateJWT(claims, 0)

	payload := map[string]string{
		"name":         user.Name,
		"username":     accountId,
		"email":        user.Email,
		"app_logo_url": configs.AppConfig().BASE_URL + "/artha-logo.png",
		"login_url":    configs.ClientConfig().ARTHA_URL + "/login",
		"action_url":   configs.ClientConfig().ARTHA_URL + "/verifyAccount?token=" + token,
	}

	apis.APIArthaSMTP().MailVerification(payload)

	return &user
}

// Me Service
/*
 * @param context *gin.Context
 * @param id string
 * @returns *models.Users
 */
func Me(context *gin.Context, id primitive.ObjectID) *models.Users {
	var user models.Users

	filter := bson.M{"_id": id}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	return &user
}
