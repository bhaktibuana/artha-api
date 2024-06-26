package helpers

import "github.com/gin-gonic/gin"

type S_HttpResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func HttpResponse(message string, httpStatus int, context *gin.Context, data interface{}) {
	response := S_HttpResponse{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
		Data:    data,
	}

	if response.Status == true {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}
