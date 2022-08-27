package controllers

import "go.mongodb.org/mongo-driver/bson/primitive"

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
}

type LearningNode struct {
	LearningNodeId          primitive.ObjectID
	LearningNodeDescription string
	Stage                   string
	NextLearningNode        []primitive.ObjectID
	PrevLearningNode        []primitive.ObjectID
	Node                    []Node
}
