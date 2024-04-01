package authRequest

import (
	"artha-api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Me Request
/*
 * @param context *gin.Context
 * @returns *string
 */
func Me(context *gin.Context) *primitive.ObjectID {
	id := context.Param("id")

	if id == ":id" {
		helpers.HttpResponse("Param 'id' is required", http.StatusUnprocessableEntity, context, nil)
		return nil
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.HttpResponse("Invalid 'id' provided", http.StatusUnprocessableEntity, context, nil)
	}

	return &_id
}
