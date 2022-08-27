package models

import (
	"context"
	"time"

	"github.com/TendonT52/tendon-api/controllers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JwtCollection struct {
	JwtCollection *mongo.Collection
}

func NewJwtCollection(jwtCollectionName string, db *DB) *JwtCollection {
	return &JwtCollection{
		JwtCollection: db.Client.Database(db.DbName).Collection(jwtCollectionName),
	}
}

func (jwtCollection *JwtCollection) GetJwtByJwtId(id primitive.ObjectID) (*controllers.Jwt, error) {
	var jwt controllers.Jwt
	filter := bson.D{{Key: "_id", Value: id}}
	err := jwtCollection.JwtCollection.FindOne(context.Background(), filter).Decode(&jwt)
	if err == mongo.ErrNoDocuments {
		return nil, IdNotFound.From(err)
	}
	if err != nil {
		return nil, UserCollectionError.From(err)
	}
	return &jwt, nil
}

func (JwtCollection *JwtCollection) UpdateJwt(jwtid primitive.ObjectID, duration time.Duration) (bool, error) {
	filter := bson.D{{Key: "_id", Value: jwtid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "expires", Value: time.Now().Add(duration)}}}}
	result, err := JwtCollection.JwtCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, ErrorWhileUpdateJwtToDatabase
	}
	return result.ModifiedCount == 1, nil
}

func (jwtCollection *JwtCollection) AddJwt(userid primitive.ObjectID, duration time.Duration) (*controllers.Jwt, error) {
	jwt := controllers.Jwt{
		UserId:  userid,
		Expires: time.Now().Add(duration),
	}
	result, err := jwtCollection.JwtCollection.InsertOne(context.Background(), jwt)
	if err != nil {
		return nil, ErrorWhileAddJwtToDatabase.From(err)
	}
	objid := result.InsertedID.(primitive.ObjectID)
	jwt.JwtId = objid
	return &jwt, nil
}

func (jwtCollection *JwtCollection) DelJwt(jwtId primitive.ObjectID) (bool, error){
	filter := bson.D{{Key: "_id", Value: jwtId}}
	number, err := jwtCollection.JwtCollection.DeleteOne(context.Background(), filter)
	if err == mongo.ErrNoDocuments {
		return false, IdNotFound.From(err)
	}
	if err != nil {
		return false, UserCollectionError.From(err)
	}
	newBool := number.DeletedCount != 0 
	return newBool, nil
}