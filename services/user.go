package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Surname       string             `json:"surname" bson:"surname"`
	Email         string             `json:"email" bson:"email"`
	Password      string             `json:"password" bson:"password"`
	Curriculum_id []int              `json:"curriculum" bson:"curriculum_id"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

type SingUpUser struct {
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SingInUser struct {
	Id           *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string              `json:"name" bson:"name"`
	Surname      string              `json:"surname" bson:"surname"`
	Email        string              `json:"email" bson:"email"`
	CurriculumId []int               `json:"curriculum" bson:"curriculum_id"`
}

func CreateSingUpUser(ctx *gin.Context) (*User, error) {
	var signIn User
	if err := ctx.ShouldBindJSON(&signIn); err != nil {
		return nil, err
	}
	signIn.Id = primitive.NewObjectID()
	hashPass := HashPassword(signIn.Password)
	signIn.Password = hashPass
	signIn.CreatedAt = time.Now()
	signIn.UpdatedAt = time.Now()
	return &signIn, nil
}
