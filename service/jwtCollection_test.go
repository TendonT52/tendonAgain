package models

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetJwtByJwtId(t *testing.T) {
	testCase := []struct {
		userId   string
		duration time.Duration
		error    error
		ok       bool
	}{
		{
			userId:   "6303306a448342f4bb47fb2e",
			duration: time.Second * 30,
			error:    nil,
			ok:       true,
		},
	}
	for _, tc := range testCase {
		t.Run("test jwt", func(t *testing.T) {
			objId, err := primitive.ObjectIDFromHex(tc.userId)
			if err != nil {
				t.Errorf("error while create object id %s", err.Error())
				return
			}
			result, err := db.JwtCollection.AddJwt(objId, tc.duration)
			if err != nil {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			jwt, err := db.JwtCollection.GetJwtByJwtId(result.JwtId)
			if err != tc.error {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			if jwt.UserId.Hex() != tc.userId {
				t.Errorf("expect %v got %v", tc.userId, jwt.UserId.Hex())
				return
			}
		})
	}
}

func TestUpdateJwt(t *testing.T) {
	testCase := []struct {
		userId         string
		duration       time.Duration
		updateDuration time.Duration
		error          error
		ok             bool
	}{
		{
			userId:         "6303306a448342f4bb47fb2e",
			duration:       time.Second * 10,
			updateDuration: time.Minute * 5,
			error:          nil,
			ok:             true,
		},
	}
	for _, tc := range testCase {
		t.Run("test jwt", func(t *testing.T) {
			objId, err := primitive.ObjectIDFromHex(tc.userId)
			if err != nil {
				t.Errorf("error while create object id %s", err.Error())
				return
			}
			result, err := db.JwtCollection.AddJwt(objId, tc.duration)
			if err != nil {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			ok, err := db.JwtCollection.UpdateJwt(result.JwtId, tc.updateDuration)
			if err != tc.error {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			if ok != tc.ok {
				t.Errorf("expect %v got %v", tc.ok, ok)
				return
			}
		})
	}
}

func TestAddJwt(t *testing.T) {
	testCase := []struct {
		userId   string
		duration time.Duration
		error    error
		ok       bool
	}{
		{
			userId:   "6303306a448342f4bb47fb2e",
			duration: time.Second * 10,
			error:    nil,
			ok:       true,
		},
	}
	for _, tc := range testCase {
		t.Run("test jwt", func(t *testing.T) {
			objId, err := primitive.ObjectIDFromHex(tc.userId)
			if err != nil {
				t.Errorf("error while create object id %s", err.Error())
				return
			}
			result, err := db.JwtCollection.AddJwt(objId, tc.duration)
			if err != tc.error {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			if !result.JwtId.IsZero() != tc.ok {
				t.Errorf("expect %v got %v", tc.ok, result.JwtId.IsZero())
				return
			}
		})
	}
}

func TestDelJwt(t *testing.T) {
	testCase := []struct {
		userId   string
		duration time.Duration
		error    error
		ok       bool
	}{
		{
			userId:   "6303306a448342f4bb47fb2e",
			duration: time.Second * 10,
			error:    nil,
			ok:       true,
		},
	}
	for _, tc := range testCase {
		t.Run("test jwt", func(t *testing.T) {
			objId, err := primitive.ObjectIDFromHex(tc.userId)
			if err != nil {
				t.Errorf("error while create object id %s", err.Error())
				return
			}
			result, err := db.JwtCollection.AddJwt(objId, tc.duration)
			if err != tc.error {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			ok, err := db.JwtCollection.DelJwt(result.JwtId)
			if err != tc.error {
				t.Errorf("expect %s got %s", tc.error, err)
				return
			}
			if ok != tc.ok {
				t.Errorf("expect %v got %v", tc.ok, ok)
				return
			}
		})
	}
}