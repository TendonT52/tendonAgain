package controllers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jwt struct {
	JwtId   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Expires time.Time          `json:"expires" bson:"expires"`
}
