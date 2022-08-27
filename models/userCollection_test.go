package models

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/TendonT52/tendon-api/controllers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUserByEmail(t *testing.T) {
	testCase := []struct {
		id            string
		name          string
		surname       string
		email         string
		errorWithCode error
	}{
		{
			id:            "6303306a448342f4bb47fb2e",
			name:          "nametest",
			surname:       "surnametest",
			email:         "test@test.com",
			errorWithCode: nil,
		},
		{
			email:         "wrong@email.com",
			errorWithCode:  EmailNotFound,
		},
	}
	for _, tc := range testCase {
		result, errWithCode := db.UserCollection.GetUserByEmail(tc.email)
		if errWithCode != nil {
			if !errors.Is(errWithCode, tc.errorWithCode) {
				t.Errorf("expect %v got %v", tc.errorWithCode, errWithCode)
				return
			}
			return 
		}
		if result.UserId.Hex() != tc.id  {
			t.Errorf("expect %s got %s", tc.id, result.UserId.Hex())
			return
		}
		if result.Firstname != tc.name {
			t.Errorf("expect %s got %s", tc.name, result.Firstname)
			return
		}
		if result.Lastname != tc.surname {
			t.Errorf("expect %s got %s", tc.surname, result.Lastname)
			return
		}

	}
}

func TestGetUserById(t *testing.T) {

	testCase := []struct {
		id            string
		name          string
		surname       string
		email         string
		errorWithCode error
	}{
		{
			id:            "6303306a448342f4bb47fb2e",
			name:          "nametest",
			surname:       "surnametest",
			email:         "test@test.com",
			errorWithCode: nil,
		},
		{
			id:            "6303306a448342f4bb47fb2a",
			name:          "nametest",
			surname:       "surnametest",
			email:         "test@test.com",
			errorWithCode: IdNotFound,
		},
	}

	for _, tc := range testCase {
		t.Run("Test Found Id", func(t *testing.T) {
			objId, err := primitive.ObjectIDFromHex(tc.id)
			if err != nil {
				t.Errorf("error while create object id %s", err.Error())
				return
			}
			result, errWithCode := db.UserCollection.GetUserById(objId)
			if errWithCode != nil {
				if !errors.Is(errWithCode, tc.errorWithCode){
					t.Errorf("expect %v got %v", tc.errorWithCode, errWithCode)
					return
				}
				return
			}

			if result.UserId.Hex() != tc.id {
				t.Errorf("expect %s got %s", tc.id, result.UserId.Hex())
				return
			}
			if result.Firstname != tc.name {
				t.Errorf("expect %s got %s", tc.name, result.Firstname)
				return
			}
			if result.Lastname != tc.surname {
				t.Errorf("expect %s got %s", tc.surname, result.Lastname)
				return
			}
			if result.Email != tc.email {
				t.Errorf("expect %s got %s", tc.email, result.Email)
				return
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	signUpUser := controllers.SignUpUser{
		Email:    fmt.Sprintf("test%d@test.com", time.Now().Unix()),
		FirstName:     fmt.Sprintf("name%d", time.Now().Unix()),
		LastName:  fmt.Sprintf("surname%d", time.Now().Unix()),
		Password: fmt.Sprintf("%d", time.Now().Unix()),
	}
	_, err := db.UserCollection.AddUser(&signUpUser)
	if err != nil {
		t.Errorf("Error while add user to db %s", err.Error())
	}
}
