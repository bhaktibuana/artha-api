package authRequest

import (
	"artha-api/src/constants"
	"artha-api/src/database"
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type S_RegisterRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"`
}

// Register Request
/*
 * @param context *gin.Context
 * @returns *S_RegisterRequest
 */
func Register(context *gin.Context) *S_RegisterRequest {
	var request S_RegisterRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	filter := bson.M{"email": request.Email}

	if count, _ := database.Users.CountDocuments(context, filter); count > 0 {
		helpers.HttpResponse(constants.EMAIL_EXISTS, http.StatusConflict, context, nil)
		return nil
	}

	return &request
}
