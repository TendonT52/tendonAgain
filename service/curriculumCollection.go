package models

import (
	"context"

	"github.com/TendonT52/tendon-api/controllers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type CurriculumCollection struct {
	CurriculumCollection *mongo.Collection
}

func NewCurriculumCollection(curriculumCollectionName string, db *DB) *CurriculumCollection {
	return &CurriculumCollection{
		CurriculumCollection: db.Client.Database(db.DbName).Collection(curriculumCollectionName),
	}
}

func (curriculumCollection *CurriculumCollection) NewCurriculum(curriculum controllers.Curriculum) (*controllers.Curriculum, error) {
	result, err := curriculumCollection.CurriculumCollection.InsertOne(context.Background(), curriculum)
	if err != nil {
		return nil, ErrorWhileAddUserToDatabase.From(err)
	}
	objid := result.InsertedID.(primitive.ObjectID)
	curriculum.CurriculumId = objid
	return &curriculum, nil
}