package authController

import (
	authRequest "artha-api/src/requests/auth"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	request := authRequest.Login(context)
	if request == nil {
		return
	}

	// user := authService.Login(context, request)
	// if user == nil {
	// 	return
	// }

	// helpers.Response("Login success", http.StatusOK, context, authResult.Login(user))
}
