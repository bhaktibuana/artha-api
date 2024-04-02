package authResult

import (
	"artha-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_VerifyEmail struct {
	ID primitive.ObjectID `json:"id"`
}

// VerifyEmail Request
/*
 * @param user *models.Users
 * @returns S_VerifyEmail
 */
func VerifyEmail(user *models.Users) S_VerifyEmail {
	return S_VerifyEmail{
		ID: user.ID,
	}
}
