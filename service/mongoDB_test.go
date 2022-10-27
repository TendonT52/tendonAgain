package models

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var db *DB

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
	db, err = NewMongoDB(os.Getenv("DB_NAME"), os.Getenv("USER_COLLECTION_NAME"), os.Getenv("JWT_COLLECTION_NAME"), os.Getenv("MONGO_DNS"))
	if err != nil {
		log.Fatalf("Error to new mongodb instant, %s", err.Error())
	}
}

func TestMongo(t *testing.T) {
	err := db.TestMongo()
	if err != nil {
		t.Error(err)
	}
}

func TestNewMongo(t *testing.T){
	if db.DbName == ""{
		t.Error("DbName is empty")
	}
	if db.Client == nil{
		t.Error("Client is nil")
	}
	if db.UserCollection == nil{
		t.Error("UserCollection is nil")
	}
}

func shutdown() {
	db.DisconnectMongo()
}
