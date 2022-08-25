package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/TendonT52/tendon-api/controllers"
)

func TestGetUserByEmail(t *testing.T) {
	testCase := []struct {
		id            string
		name          string
		surname       string
		email         string
		errorWithCode errorWithCode
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
			errorWithCode: &EmailNotFound,
		},
	}
	for _, tc := range testCase {
		result, errWithCode := db.UserCollection.GetUserByEmail(tc.email)
		if errWithCode != nil {
			if errWithCode.GetKind() != tc.errorWithCode.GetKind() {
				t.Errorf("expect %v got %v", tc.errorWithCode, errWithCode)
				return
			}
			return 
		}
		if result.Id.Hex() != tc.id  {
			t.Errorf("expect %s got %s", tc.id, result.Id.Hex())
			return
		}
		if result.Name != tc.name {
			t.Errorf("expect %s got %s", tc.name, result.Name)
			return
		}
		if result.Surname != tc.surname {
			t.Errorf("expect %s got %s", tc.surname, result.Surname)
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
		errorWithCode errorWithCode
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
			errorWithCode: &IdNotFound,
		},
	}

	for _, tc := range testCase {
		t.Run("Test Found Id", func(t *testing.T) {
			result, errWithCode := db.UserCollection.GetUserById(tc.id)
			if errWithCode != nil {
				if errWithCode.GetKind() != tc.errorWithCode.GetKind() {
					t.Errorf("expect %v got %v", tc.errorWithCode, errWithCode)
					return
				}
				return
			}

			if result.Id.Hex() != tc.id {
				t.Errorf("expect %s got %s", tc.id, result.Id.Hex())
				return
			}
			if result.Name != tc.name {
				t.Errorf("expect %s got %s", tc.name, result.Name)
				return
			}
			if result.Surname != tc.surname {
				t.Errorf("expect %s got %s", tc.surname, result.Surname)
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
		Name:     fmt.Sprintf("name%d", time.Now().Unix()),
		Surname:  fmt.Sprintf("surname%d", time.Now().Unix()),
		Password: fmt.Sprintf("%d", time.Now().Unix()),
	}
	_, err := db.UserCollection.AddUser(&signUpUser)
	if err != nil {
		t.Errorf("Error while add user to db %s", err.Error())
	}
}
