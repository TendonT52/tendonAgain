package models

import (
	"context"

	"github.com/TendonT52/tendon-api/controllers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	Collection *mongo.Collection
}

func NewUserCollection(userCollectionName string, db *DB) *UserCollection {
	return &UserCollection{
		Collection: db.Client.Database(db.DbName).Collection(userCollectionName),
	}
}

func (userCollection *UserCollection) GetUserByEmail(user controllers.HaveEmail) (*controllers.SignInUser, errorWithCode) {
	var signInuser controllers.SignInUser
	email := bson.D{{Key: "email", Value: user.GetEmail()}}
	err := userCollection.Collection.FindOne(context.Background(), email).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, EmailNotFound.New()
	}
	if err != nil {
		return nil, DbError.From(err)
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) GetUserById(user *controllers.SignInUser) (*controllers.SignInUser, errorWithCode) {
	var signInuser controllers.SignInUser
	err := userCollection.Collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: user.Id}}).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, IdNotFound.New()
	}
	if err != nil {
		return nil, DbError.New()
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) AddUser(user *controllers.SignUpUser) (*controllers.SignInUser, errorWithCode) {
	_, errWithCode := userCollection.GetUserByEmail(user)
	if errWithCode == nil {
		return nil, UserIsAlreadyExist.New()
	}
	result, err := userCollection.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, ErrorWhileAddUserToDatabase.New()
	}
	objid := result.InsertedID.(primitive.ObjectID)
	singInUser := controllers.SignInUser{
		Id: &objid,
	}
	return &singInUser, nil
}