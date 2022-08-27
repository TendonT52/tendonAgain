package handlers

import (
	"net/http"
	"time"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/controllers"
	"github.com/TendonT52/tendon-api/error"
	"github.com/TendonT52/tendon-api/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	app *configs.App
}

func NewHandlerAuth(app *configs.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (handleAuth *AuthHandler) HandleSignUp(ctx *gin.Context) {
	signUpUser := controllers.SignUpUser{}
	err := ctx.BindJSON(&signUpUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "request in wrong format"})
		return
	}
	signUpUser.Password = services.HashPassword(signUpUser.Password)
	_, err = handleAuth.app.MongoDB.UserCollection.AddUser(&signUpUser)
	if err != nil {
		ctx.JSON(err.(error.ErrorWithCode).Code, err.(error.ErrorWithCode).Response)
		return
	}
	ctx.JSON(
		http.StatusAccepted,
		gin.H{
			"name":    signUpUser.FirstName,
			"surname": signUpUser.LastName,
			"email":   signUpUser.Email,
			"token":   handleAuth.app.JwtSecret.GenerateAccessToken("", "",time.Minute*15),
		},
	)
}

func (handleAuth *AuthHandler) HandleSignIn(ctx *gin.Context) {
	signUpUser := controllers.SignInUser{}
	err := ctx.BindJSON(&signUpUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "request in wrong format"})
		return
	}
	signInUser, err := handleAuth.app.MongoDB.UserCollection.GetUserByEmail(signUpUser.Email)
	if err != nil {
			ctx.JSON(http.StatusNotFound, err.(error.ErrorWithCode).Response)
			return
	}
	if !services.CheckPasswordHash(signUpUser.Password, signInUser.Password) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "email or password is incorrect"})
		return
	}
	ctx.JSON(
		http.StatusAccepted,
		gin.H{
			"name":    signInUser.Firstname,
			"surname": signInUser.Lastname,
			"email":   signInUser.Email,
			"token":   handleAuth.app.JwtSecret.GenerateAccessToken("","",  time.Minute*15),
		},
	)
}
