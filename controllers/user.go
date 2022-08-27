package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Admin = "admin"
	Teacher = "teacher"
	User = "student"
)

type SignUpUser struct {
	FirstName     string `json:"name" bson:"name"`
	LastName  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SignInUser struct {
	UserId        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname     string             `json:"name" bson:"name"`
	Lastname      string             `json:"surname" bson:"surname"`
	Email         string             `json:"email" bson:"email"`
	Password      string             `json:"password" bson:"password"`
	Role          string             `json:"role" bson:"role"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	SubCurriculum []SubCurriculum    `json:"subcurriculum" bson:"subcurriculum"`
}

type SubCurriculum struct {
	CurriculumId          primitive.ObjectID
	SubCurriculum         string
	CurriculumDescription string
	ProGress              int
}
