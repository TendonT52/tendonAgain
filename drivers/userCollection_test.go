package drivers

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/TendonT52/tendon-api/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db *DB;

func TestMain(m *testing.M) {
	db = setup()
	code := m.Run()
	shutdown(db)
	os.Exit(code)
}

func setup() ( *DB ){
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
	db := NewMongoDB(os.Getenv("DB_NAME"), os.Getenv("MONGO_DNS"))
	return	db
}

func shutdown(db *DB) {
	db.DisconnectMongo()
}

func TestNewUserCollection(t *testing.T) {
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
}

func TestIsUserExitNotFound(t *testing.T){
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
	result := userCollection.IsUserExit(services.User{Email : "test@gmail.coma"})
	if result != false {
		t.Error("IsUserExit got wrong")
	}
}

func TestIsUserExitFound(t *testing.T){
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
	result := userCollection.IsUserExit(services.User{Email : "test@gmail.com"})
	if result != true {
		t.Error("IsUserExit got wrong")
	}
}

func TestAddUser(t *testing.T){
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
	var signIn services.User
	signIn.Id = primitive.NewObjectID()
	signIn.Email = fmt.Sprintf("emailTest %d", time.Now().Unix())
	signIn.Name = fmt.Sprintf("nameTest %d", time.Now().Unix())
	signIn.Surname = fmt.Sprintf("surnameTest %d", time.Now().Unix())
	signIn.Password = services.HashPassword("1234")
	signIn.CreatedAt = time.Now()
	signIn.UpdatedAt = time.Now()

	result, ok := userCollection.AddUser(signIn)
	if !ok {
		t.Error("Wrong ans")
	}
	log.Printf("result, %+v", result)
}

func TestGetUserFound(t *testing.T) {
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
	var signIn services.User
	signIn.Email = "emailTest 1660475774"
	user, ok := userCollection.GetUser(signIn)
	if !ok {
		t.Error("Wrong ans")
	}
	log.Printf("result, %+v", user)
}

func TestGetUserNotFound(t *testing.T) {
	userCollection := NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), db)	
	if userCollection.Collection.Name() != os.Getenv("USER_COLLECTION_NAME") {
		t.Errorf("Error while get collection from db");	
	}
	var signIn services.User
	signIn.Email = "emailThaNotFound"
	user, ok := userCollection.GetUser(signIn)
	if ok {
		t.Error("Wrong ans")
	}
	log.Printf("result, %+v", user)
}