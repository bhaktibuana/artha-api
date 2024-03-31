package authResult

import (
	"artha-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginResult struct {
	ID      primitive.ObjectID `json:"id"`
	ArthaId string             `json:"artha_id"`
	Token   string             `json:"token"`
}

func Login(user *models.Users) LoginResult {
	return LoginResult{
		ID:      user.ID,
		ArthaId: user.Username + "#" + user.Tag,
	}
}
