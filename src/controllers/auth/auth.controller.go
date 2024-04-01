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
	user := authService.Me(context)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, authResult.Me(user))
}

// LoginVerify2FA Controller
/*
 * @param context *gin.Context
 * @returns
 */
func LoginVerify2FA(context *gin.Context) {
	id, request := authRequest.LoginVerify2FA(context)
	if request == nil {
		return
	}

	user := authService.LoginVerify2FA(context, *id, request)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, authResult.Verify2FA(user))
}

// Verify2FA Controller
/*
 * @param context *gin.Context
 * @returns
 */
func Verify2FA(context *gin.Context) {
	request := authRequest.Verify2FA(context)
	if request == nil {
		return
	}

	user := authService.Verify2FA(context, request)
	if user == nil {
		return
	}

	helpers.HttpResponse(constants.REQUEST_SUCCESS, http.StatusOK, context, authResult.Verify2FA(user))
}
