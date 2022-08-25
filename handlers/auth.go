package handlers

import (
	"net/http"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/controllers"
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
	signInUser, errWithCode := handleAuth.app.MongoDB.UserCollection.AddUser(&signUpUser)
	if errWithCode != nil {
		ctx.JSON(errWithCode.GetCode(), errWithCode.GetValue())
		return
	}
	ctx.JSON(
		http.StatusAccepted,
		gin.H{
			"name":    signUpUser.Name,
			"surname": signUpUser.Surname,
			"email":   signUpUser.Email,
			"token":   handleAuth.app.JwtSecret.GenerateAccessToken(*signInUser),
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
	signInUser, errWithCode := handleAuth.app.MongoDB.UserCollection.GetUserByEmail(signUpUser.Email)
	if errWithCode != nil {
			ctx.JSON(http.StatusNotFound, errWithCode.GetValue())
			return
	}
	if !services.CheckPasswordHash(signUpUser.Password, signInUser.Password) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": "email or password is incorrect"})
		return
	}
	ctx.JSON(
		http.StatusAccepted,
		gin.H{
			"name":    signInUser.Name,
			"surname": signInUser.Surname,
			"email":   signInUser.Email,
			"token":   handleAuth.app.JwtSecret.GenerateAccessToken(*signInUser),
		},
	)
}
