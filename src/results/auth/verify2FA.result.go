package authResult

import (
	"artha-api/src/helpers"
	"artha-api/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_Verify2FAResult struct {
	ID      primitive.ObjectID `json:"id"`
	ArthaId string             `json:"artha_id"`
	Token   string             `json:"token"`
}

// Verify2FA Request
/*
 * @param user *models.Users
 * @returns S_Verify2FAResult
 */
func Verify2FA(user *models.Users) S_Verify2FAResult {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"tag":      user.Tag,
		"username": user.Username,
		"name":     user.Name,
	}

	token, _ := helpers.GenerateJWT(claims, time.Hour*24*30)

	return S_Verify2FAResult{
		ID:      user.ID,
		ArthaId: user.Username + "#" + user.Tag,
		Token:   token,
	}
}
