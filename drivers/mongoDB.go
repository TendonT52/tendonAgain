package drivers

import (
	"context"
	"log"
	"time"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	DbName string
	Client *mongo.Client
	UserCollection *UserCollection
}

func NewMongoDB(dbName string, dsn string) *DB {
	client := ConnectMongo(dsn)
	db := DB{
		DbName: dbName,
		Client: client,
	}
	db.TestMongo()
	return &db
}

func ConnectMongo(dsn string) *mongo.Client {
	log.Println("Connecting to database...")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("Error while connecting %s", err)
	}
	log.Println("Connected")
	return client
}

func (client *DB)DisconnectMongo(){
	log.Println("Disconnecting to database...")
	if err := client.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error while disconnect %s", err)
		}
	log.Println("Disconnected")
}

func (client *DB)TestMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := client.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error while ping %s", err)
	}
	log.Printf("Test mongo success")	
}

