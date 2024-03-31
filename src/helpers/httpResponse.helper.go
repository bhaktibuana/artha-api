package helpers

import "github.com/gin-gonic/gin"

type S_HttpResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

// HttpResponse Helper
/*
 * @param message string
 * @param httpStatus int
 * @param context *gin.Context
 * @param data interface{}
 * @returns
 */
func HttpResponse(message string, httpStatus int, context *gin.Context, data interface{}) {
	response := S_HttpResponse{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
		Data:    data,
	}

	if response.Status {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}
