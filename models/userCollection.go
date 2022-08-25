package models

import (
	"context"
	"time"

	"github.com/TendonT52/tendon-api/controllers"
	"github.com/TendonT52/tendon-api/services"
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

func (userCollection *UserCollection) GetUserByEmail(email string) (*controllers.SignInUser, errorWithCode) {
	var signInuser controllers.SignInUser
	filter := bson.D{{Key: "email", Value: email}}
	err := userCollection.Collection.FindOne(context.Background(), filter).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, EmailNotFound.New()
	}
	if err != nil {
		return nil, DbError.From(err)
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) GetUserById(id string) (*controllers.SignInUser, errorWithCode) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, IdNotValid.New()
	}
	var signInuser controllers.SignInUser
	filter := bson.D{{Key: "_id", Value: objId}}
	err = userCollection.Collection.FindOne(context.Background(), filter).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, IdNotFound.New()
	}
	if err != nil {
		return nil, DbError.New()
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) AddUser(signUpUser *controllers.SignUpUser) (*controllers.SignInUser, errorWithCode) {
	_, errWithCode := userCollection.GetUserByEmail(signUpUser.Email)
	if errWithCode == nil {
		return nil, UserIsAlreadyExist.New()
	}

	signInUser := controllers.SignInUser{}
	signInUser.Name = signUpUser.Name
	signInUser.Surname = signUpUser.Surname
	signInUser.Email = signUpUser.Email
	signInUser.Password = services.HashPassword(signUpUser.Password)
	signInUser.CurriculumId = []primitive.ObjectID{}
	signInUser.CreatedAt = time.Now()
	signInUser.UpdatedAt = time.Now()

	result, err := userCollection.Collection.InsertOne(context.Background(), signInUser)
	if err != nil {
		return nil, ErrorWhileAddUserToDatabase.New()
	}
	objid := result.InsertedID.(primitive.ObjectID)
	signInUser.Id = objid
	return &signInUser, nil
}
