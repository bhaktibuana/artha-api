package authResult

import (
	"artha-api/src/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_MeResult struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Tag         string             `json:"tag"`
	AccountId   string             `json:"account_id"`
	Email       string             `json:"email"`
	AccountType string             `json:"account_type"`
	Status      string             `json:"status"`
	Photo       string             `json:"photo"`
}

// Me Request
/*
 * @param user *models.Users
 * @returns S_MeResult
 */
func Me(user *models.Users) S_MeResult {
	return S_MeResult{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Tag:         user.Tag,
		AccountId:   fmt.Sprintf("%s#%s", user.Username, user.Tag),
		Email:       user.Email,
		AccountType: user.AccountType,
		Status:      user.Status,
		Photo:       user.Photo,
	}
}
