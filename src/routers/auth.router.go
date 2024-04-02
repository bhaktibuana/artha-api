package routers

import (
	authController "artha-api/src/controllers/auth"
	routersTemplate "artha-api/src/routers/templates"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	basePath := "/auth"
	publicRoute := routersTemplate.NewPublicRoute(basePath, router)
	privateRoute := routersTemplate.NewPrivateRoute(basePath, router)

	publicRoute.POST("/login", authController.Login)
	publicRoute.POST("/register", authController.Register)
	privateRoute.GET("/me", authController.Me)
	publicRoute.POST("/:id/login-verify-2fa", authController.LoginVerify2FA)
	privateRoute.POST("/verify-2fa", authController.Verify2FA)
	privateRoute.GET("/get-new-2fa-uri", authController.GetNew2FAUri)
	privateRoute.PUT("/update-2fa-secret", authController.Update2FASecret)
	publicRoute.POST("/verify-email", authController.VerifyEmail)
}
