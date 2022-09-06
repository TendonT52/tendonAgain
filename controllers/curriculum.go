package controllers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PublicAccess    = "public"
	ProtectedAccess = "protected"
	PrivateAccess   = "private"
)

type Curriculum struct {
	CurriculumId    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CurriculumName  string             `json:"curriculumName " bson:"curriculum_name"`
	SubLearningNode []SubLearningNode  `json:"subLearningNode" bson:"sub_learning_node"`
	Access          string             `json:"access" bson:"access"`
	CreateBy        primitive.ObjectID `json:"createBy" bson:"create_by"`
	CreatedAt       time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updated_at"`
}

type SubLearningNode struct {
	LearningNodeId          primitive.ObjectID   `json:"learningNodeId" bson:"learning_node_id"`
	LearningNodeDescription string               `json:"learningNodeDescription" bson:"learning_node_description"`
	NextLearningNode        []primitive.ObjectID `json:"nextLearningNode" bson:"next_learning_node"`
	PrevLearningNode        []primitive.ObjectID `json:"prevLearningNode" bson:"prev_learning_node"`
}
