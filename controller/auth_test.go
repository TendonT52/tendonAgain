package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app *configs.App

type Response struct {
	Name    string
	Surname string
	Email   string
	Token   string
}

type Message struct {
	Message    string
}

func TestMain(m *testing.M) {
	setUp()
	app = configs.NewApp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app = configs.NewApp()
}

func setUpAuthHandler() (*gin.Engine, *AuthHandler) {
	r := gin.Default()
	authHandler := NewHandlerAuth(app)
	return r, authHandler
}

func TestHandlerSignUp(t *testing.T) {
	email := fmt.Sprintf("test%d@test.com", time.Now().Unix())
	name :=fmt.Sprintf("name%d", time.Now().Unix())
	surname :=fmt.Sprintf("surname%d", time.Now().Unix())
	password :=fmt.Sprintf("%d", time.Now().Unix())
	r, authHandler := setUpAuthHandler()
	r.POST("/signup", authHandler.HandleSignUp)
	signUpUser := controllers.SignUpUser{
		Email:   email,
		FirstName:    name, 
		LastName:  surname,
		Password:  password,
	}
	jsonValue, _ := json.Marshal(signUpUser)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Error("Can't unmarshal json response body")
	}
	if res.Email != email {
		t.Errorf("expect %s but got %s", email, res.Email)
	}
	if res.Name != name {
		t.Errorf("expect %s but got %s", name, res.Name)
	}
	if res.Surname != surname {
		t.Errorf("expect %s but got %s", surname, res.Surname)
	}
	_, err = authHandler.app.JwtSecret.ValidateToken(res.Token)
	if err != nil {
		t.Errorf("Token can't validate %s", err.Error())
	}

	reqF, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	reqF.Header.Add("Content-Type", "application/json")
	wF := httptest.NewRecorder()
	r.ServeHTTP(wF, reqF)
	resF := Message{}
	err = json.Unmarshal(wF.Body.Bytes(), &resF)
	if err != nil {
		t.Error("Can't unmarshal json response body")
	}
	if resF.Message != "This email already in use"{
		t.Errorf("expect This email already in use but got %s", resF.Message)
	}
}

func TestHandleSignIn(t *testing.T) {
	email := "test@test.com"
	password := "test"
	r, authHandler := setUpAuthHandler()
	r.POST("/signin", authHandler.HandleSignIn)
	signUpUser := controllers.SignUpUser{
		Email:    email,
		Password: password,
	}
	jsonValue, _ := json.Marshal(signUpUser)
	req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
		res := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Error("Can't unmarshal json response body")
	}
	if res.Email != email {
		t.Errorf("expect %s but got %s", email, res.Email)
	}
	_, err = authHandler.app.JwtSecret.ValidateToken(res.Token)
	if err != nil {
		t.Errorf("Token can't validate %s", err.Error())
	}
}
