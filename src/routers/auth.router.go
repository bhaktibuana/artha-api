package routers

import (
	authController "artha-api/src/controllers/auth"
	"artha-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	basePath := "/auth"
	authGroup := router.Group(basePath)
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
		authGroup.GET("/:id", middlewares.IsAuthenticate, authController.Me)
	}
}
