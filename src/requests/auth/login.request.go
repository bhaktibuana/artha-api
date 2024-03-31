package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S_LoginRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"`
}

func Login(context *gin.Context) *S_LoginRequest {
	var request S_LoginRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.HttpResponse(err.Error(), http.StatusBadRequest, context, nil)
		return nil
	}

	return &request
}
