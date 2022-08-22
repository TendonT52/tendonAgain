package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/TendonT52/tendon-api/controllers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUserByEmail(t *testing.T) {
	emailFound := "test@test.com"
	emailNotFound := "Test@test.com"
	signUpUserFound := controllers.SignUpUser{
			Email: emailFound,
	}
	result, err := db.UserCollection.GetUserByEmail(&signUpUserFound)
	if err != nil {
		t.Errorf("User should found But not found \n %s", err.Error())
		return
	}
	if result.Email != emailFound{
		t.Errorf("User found but incorrect \n %s", err.Error())
		return
	}
	signUpUserNotFound := controllers.SignUpUser{
			Email: emailNotFound,
	}
	_, err = db.UserCollection.GetUserByEmail(&signUpUserNotFound)
	if err.GetKind() != EmailNotFound.GetKind() {
		t.Errorf("User should not found but found \n %s", err.Error())
		return
	}
}

func TestGetUserById(t *testing.T) {
	objFound := "6303306a448342f4bb47fb2e"
	objNotFound := "6303306a448342f4bb47fb2a"
	objId, err := primitive.ObjectIDFromHex(objFound)
	if err != nil {
		t.Errorf("Can't create object id from given hex string")
		return
	}
	signInUser := controllers.SignInUser{
		Id: &objId,
	}
	result, errWithCode := db.UserCollection.GetUserById(&signInUser)
	if errWithCode != nil {
		t.Errorf("User should found But not found \n %s", err.Error())
		return
	}
	if result.Id.Hex() != objFound{
		t.Errorf("User found but incorrect \n %s", err.Error())
		return
	}
	objId, err = primitive.ObjectIDFromHex(objNotFound)
	if err != nil {
		t.Errorf("Can't create object id from given hex string")
	}
	signInUser = controllers.SignInUser{
		Id: &objId,
	}
	_, errWithCode = db.UserCollection.GetUserById(&signInUser)
	if errWithCode.GetKind() != IdNotFound.GetKind(){
		t.Error("User should not found but found \n")
	}
}

func TestAddUser(t *testing.T) {
	signUpUser := controllers.SignUpUser{
		Email:    fmt.Sprintf("test%d@test.com", time.Now().Unix()),
		Name:     fmt.Sprintf("name%d", time.Now().Unix()),
		Surname:  fmt.Sprintf("surname%d", time.Now().Unix()),
		Password: fmt.Sprintf("%d", time.Now().Unix()),
	}
	_, err :=db.UserCollection.AddUser(&signUpUser)
	if err != nil {
		t.Errorf("Error while add user to db %s", err.Error())
	}
}
