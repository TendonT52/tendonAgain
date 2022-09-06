package controllers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Required   = "required"
	ShouldKnow = "shouldknow"
	Optional   = "optional"
)

const (
	Text = "text"
	Image = "image"
	Video = "video"
)

type Node struct {
	Type string
	Name string
	Data string
	CreatedAt       time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updated_at"`
}

type LearningNode struct {
	LearningNodeId          primitive.ObjectID
	LearningNodeDescription string
	Stage                   string
	NextLearningNode        []primitive.ObjectID
	PrevLearningNode        []primitive.ObjectID
	Node                    []Node
}
