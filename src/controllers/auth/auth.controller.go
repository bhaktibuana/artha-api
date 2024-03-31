package authController

import (
	"artha-api/src/helpers"
	authRequest "artha-api/src/requests/auth"
	authResult "artha-api/src/results/auth"
	authService "artha-api/src/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	request := authRequest.Login(context)
	if request == nil {
		return
	}

	user := authService.Login(context, request)
	if user == nil {
		return
	}

	helpers.HttpResponse("Login success", http.StatusOK, context, authResult.Login(user))
}
