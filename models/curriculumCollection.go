package models

import "go.mongodb.org/mongo-driver/mongo"


type CurriculumCollection struct {
	CurriculumCollection *mongo.Collection
}

func NewCurriculumCollection(curriculumCollectionName string, db *DB) *CurriculumCollection {
	return &CurriculumCollection{
		CurriculumCollection: db.Client.Database(db.DbName).Collection(curriculumCollectionName),
	}
}
