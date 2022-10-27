package models

import (
	"context"
	"time"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	DbName string
	Client *mongo.Client
	UserCollection *UserCollection
	JwtCollection *JwtCollection
}

func NewMongoDB(dbName string, userCollectionName string, jwtCollectionName string,dsn string) (*DB, error) {
	client, err := ConnectMongo(dsn)
	if err != nil {
		return nil, err
	}
	db := DB{
		DbName: dbName,
		Client: client,
	}
	db.UserCollection = NewUserCollection(userCollectionName, &db)
	db.JwtCollection = NewJwtCollection(jwtCollectionName, &db)
	return &db, nil
}

func ConnectMongo(dsn string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *DB)DisconnectMongo() error {
	if err := client.Client.Disconnect(context.TODO()); err != nil {
			return err
		}
	return nil
}

func (client *DB)TestMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := client.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

