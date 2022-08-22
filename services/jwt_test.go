package services

import (
	"log"
	"os"
	"testing"

	"github.com/TendonT52/tendon-api/controllers"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestGenerateToken(t *testing.T) {
	id := "62fceec0125e83372bcd6cce"
	email := "test1660743360@test.com"
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Can't create object id from given hex string")
		return
	}
	singInUser := controllers.SignInUser{
		Id:           &objId,
		Email:        email,
		CurriculumId: []int{},
	}
	tokenString := jwtServices.GenerateAccessToken(singInUser)
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtServices.secretKey), nil
	})
	if err != nil {
		t.Errorf("Error while decode jwt token, %s", err)
		return
	}
	if claims["email"] != email && claims["id"] != id {
		t.Error("jwt token give wrong email or id")
		return
	}
}

func TestValidateToken(t *testing.T) {
	id := "62fceec0125e83372bcd6cce"
	email := "test1660743360@test.com"
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Can't create object id from given hex string")
		return
	}
	singInUser := controllers.SignInUser{
		Id:           &objId,
		Email:        email,
		CurriculumId: []int{},
	}
	tokenString := jwtServices.GenerateAccessToken(singInUser)
	token, err := jwtServices.ValidateToken(tokenString)
	if err != nil {
		t.Error("error while validate token")
		return
	}
	if !token.Valid {
		t.Error("token should valid but not")
		return
	}
	tokenString = "asdf.asdf"
	token, err = jwtServices.ValidateToken(tokenString)
	if err != nil {
		return
	}
	if token.Valid {
		t.Error("token should not valid")
	}
}
