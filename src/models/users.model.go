package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Username    string             `bson:"username" json:"username"`
	Tag         string             `bson:"tag" json:"tag"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	Birthdate   time.Time          `bson:"birthdate" json:"birthdate"`
	Gender      string             `bson:"gender" json:"gender"`
	AccountType string             `bson:"account_type" json:"account_type"`
	Photo       string             `bson:"photo" json:"photo"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}
