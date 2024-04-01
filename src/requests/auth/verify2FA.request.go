package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_Verify2FA struct {
	OTP string `json:"otp" binding:"required"`
}

// Verify2FA Request
/*
 * @param context *gin.Context
 * @returns *S_Verify2FA
 */
func Verify2FA(context *gin.Context) *S_Verify2FA {
	var request S_Verify2FA

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	return &request
}
