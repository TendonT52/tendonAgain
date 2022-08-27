package services

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/joho/godotenv"
)

var jwtServices *JwtServices

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
	jwtServices = NewJwtServices(os.Getenv("JWT_SECRET"))
}

func TestGenerateAccessToken(t *testing.T) {
	testCase := []struct {
		id       string
		duration time.Duration
	}{
		{
			id:       "62fceec0125e83372bcd6cce",
			duration: time.Minute * 1,
		},
	}

	for _, tc := range testCase {
		tokenString := jwtServices.GenerateAccessToken(tc.id, "", tc.duration)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtServices.secretKey), nil
		})
		if err != nil {
			t.Errorf("Error while decode jwt token, %s", err)
			return
		}
		if claims["id"] != tc.id {
			t.Error("jwt token give wrong email or id")
			return
		}
	}
}

func TestValidateToken(t *testing.T) {
	testCast := []struct {
		id       string
		email    string
		duration time.Duration
		valid    bool
		err      error
		expectId string
	}{
		{
			id:       "62fceec0125e83372bcd6cce",
			duration: time.Minute,
			err:      nil,
			expectId: "62fceec0125e83372bcd6cce",
		},
		{
			id:       "62fceec0125e83372bcd6cce",
			duration: 0,
			err:      TokenExpired,
			expectId: "",
		},
	}
	for _, tc := range testCast {
		tokenString := jwtServices.GenerateAccessToken(tc.id, "", tc.duration)
		id, err := jwtServices.ValidateToken(tokenString)
		if !errors.Is(err, tc.err) {
			t.Errorf("expect %s got %s", tc.err, err)
			return
		}
		if tc.expectId != id {
			t.Errorf("expect %s got %s", tc.expectId, id)
			return
		}
	}
}
