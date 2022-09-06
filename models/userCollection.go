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
	UserCollection *mongo.Collection
}

func NewUserCollection(userCollectionName string, db *DB) *UserCollection {
	return &UserCollection{
		UserCollection: db.Client.Database(db.DbName).Collection(userCollectionName),
	}
}

func (userCollection *UserCollection) GetUserByEmail(email string) (*controllers.SignInUser, error) {
	var signInuser controllers.SignInUser
	filter := bson.D{{Key: "email", Value: email}}
	err := userCollection.UserCollection.FindOne(context.Background(), filter).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, EmailNotFound.From(err)
	}
	if err != nil {
		return nil, UserCollectionError.From(err)
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) GetUserById(objId primitive.ObjectID) (*controllers.SignInUser, error) {
	var signInuser controllers.SignInUser
	filter := bson.D{{Key: "_id", Value: objId}}
	err := userCollection.UserCollection.FindOne(context.Background(), filter).Decode(&signInuser)
	if err == mongo.ErrNoDocuments {
		return nil, IdNotFound.From(err)
	}
	if err != nil {
		return nil, UserCollectionError.From(err)
	}
	return &signInuser, nil
}

func (userCollection *UserCollection) DelUserById(objId primitive.ObjectID) (bool, error){
	filter := bson.D{{Key: "_id", Value: objId}}
	number, err := userCollection.UserCollection.DeleteOne(context.Background(), filter)
	if err == mongo.ErrNoDocuments {
		return false, IdNotFound.From(err)
	}
	if err != nil {
		return false, UserCollectionError.From(err)
	}
	newBool := number.DeletedCount != 0 
	return newBool, nil
}


func (userCollection *UserCollection) AddUser(signUpUser *controllers.SignUpUser) (*controllers.SignInUser, error) {
	_, errWithCode := userCollection.GetUserByEmail(signUpUser.Email)
	if errWithCode == nil {
		return nil, UserIsAlreadyExist.New()
	}

	signInUser := controllers.SignInUser{}
	signInUser.FirstName = signUpUser.FirstName
	signInUser.LastName = signUpUser.LastName
	signInUser.Email = signUpUser.Email
	signInUser.Password = services.HashPassword(signUpUser.Password)
	signInUser.SubCurriculum = []controllers.SubCurriculum{}
	signInUser.CreatedAt = time.Now()
	signInUser.UpdatedAt = time.Now()

	result, err := userCollection.UserCollection.InsertOne(context.Background(), signInUser)
	if err != nil {
		return nil, ErrorWhileAddUserToDatabase.From(err)
	}
	objid := result.InsertedID.(primitive.ObjectID)
	signInUser.UserId = objid
	return &signInUser, nil
}
