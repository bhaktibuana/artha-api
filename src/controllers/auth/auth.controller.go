package authController

import (
	"artha-api/src/constants"
	"artha-api/src/helpers"
	authRequest "artha-api/src/requests/auth"
	authResult "artha-api/src/results/auth"
	authService "artha-api/src/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Login(context *gin.Context) {
	request := authRequest.Login(context)
	if request == nil {
		return
	}

	user := authService.Login(context, request)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.LOGIN_SUCCESS, http.StatusOK, context, authResult.Login(user))
}

// Register Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Register(context *gin.Context) {
	request := authRequest.Register(context)
	if request == nil {
		return
	}

	user := authService.Register(context, request)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.REGISTER_SUCCESS, http.StatusOK, context, authResult.Register(user))
}

// Me Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Me(context *gin.Context) {
	id := authRequest.Me(context)
	if id == nil {
		return
	}

	user := authService.Me(context, *id)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, authResult.Me(user))
}
