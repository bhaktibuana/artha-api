package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_LoginVerify2FA struct {
	OTP string `json:"otp" binding:"required"`
}

// LoginVerify2FA Request
/*
 * @param context *gin.Context
 * @returns (*string, *S_LoginVerify2FA) (id, requestBody)
 */
func LoginVerify2FA(context *gin.Context) (*string, *S_LoginVerify2FA) {
	var request S_LoginVerify2FA

	id := context.Param("id")

	if id == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil, nil
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusUnprocessableEntity, context, nil)
		return nil, nil
	}

	return &id, &request
}
