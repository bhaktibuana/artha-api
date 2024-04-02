package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_VerifyEmail struct {
	Token string `json:"token" binding:"required"`
}

// VerifyEmail Request
/*
 * @param context *gin.Context
 * @returns *S_VerifyEmail
 */
func VerifyEmail(context *gin.Context) *S_VerifyEmail {
	var request S_VerifyEmail

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	return &request
}
