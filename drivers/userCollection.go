package drivers

import (
	"context"
	"fmt"

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

func (userCollection *UserCollection) IsUserExit(user services.User) bool {
	a := userCollection.Collection.FindOne(context.Background(), bson.D{{Key: "email", Value: user.Email}})
	var result bson.D
	a.Decode(&result)
	if result == nil {
		return false
	} else {
		return true
	}
}

func (userCollection *UserCollection) AddUser(user services.User) (services.User, bool) {
	if userCollection.IsUserExit(user) {
		return user, false
	}
	result, err := userCollection.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return user, false
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, true
}

func (userCollection *UserCollection) GetUser(user services.User) (services.User, bool) {
	doc := userCollection.Collection.FindOne(context.Background(), bson.D{{Key: "email", Value: user.Email}})
	err := doc.Decode(&user)
	fmt.Printf("result, %+v", user)
	if err == mongo.ErrNoDocuments {
		return user, false
	}
	return user, true
}
