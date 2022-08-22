package controllers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignInUser struct {
	Name         string              `json:"name" bson:"name"`
	Surname      string              `json:"surname" bson:"surname"`
	Email        string              `json:"email" bson:"email"`
	Id           *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Password     string              `json:"password" bson:"password"`
	CurriculumId []int               `json:"curriculum" bson:"curriculum_id"`
	CreatedAt    time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at" bson:"updated_at"`
}

func (signInUser *SignInUser) GetEmail() string {
	return signInUser.Email
}
