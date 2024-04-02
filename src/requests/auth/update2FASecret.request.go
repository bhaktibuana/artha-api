package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_Update2FASecret struct {
	OTP       string `json:"otp" binding:"required"`
	SecretKey string `json:"secret_key" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"`
}

// Update2FASecret Request
/*
 * @param context *gin.Context
 * @returns *S_Update2FASecret
 */
func Update2FASecret(context *gin.Context) *S_Update2FASecret {
	var request S_Update2FASecret

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}
	return &request
}
