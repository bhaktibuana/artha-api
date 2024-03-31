package authResult

import (
	"artha-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_RegisterResult struct {
	ID primitive.ObjectID `json:"id"`
}

// Register Request
/*
 * @param user *models.Users
 * @returns S_RegisterResult
 */
func Register(user *models.Users) S_RegisterResult {
	return S_RegisterResult{
		ID: user.ID,
	}
}
