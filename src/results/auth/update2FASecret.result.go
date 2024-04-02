package authResult

import (
	"artha-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_Update2FASecret struct {
	ID primitive.ObjectID `json:"id"`
}

// Update2FASecret Request
/*
 * @param user *models.Users
 * @returns S_Update2FASecret
 */
func Update2FASecret(user *models.Users) S_Update2FASecret {
	return S_Update2FASecret{
		ID: user.ID,
	}
}
