package handlers

import (
	"net/http"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/services"
	"github.com/gin-gonic/gin"
)

var app *configs.App

func InitHadler(ap *configs.App) {
	app = ap
}

func HandleSignUp(ctx *gin.Context) {
	user, err := services.CreateSingUpUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user_found, ok := App.MongoDB.UserCollection.AddUser(*user)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email already be use"})
		return
	}
	token := app.JwtSecret.GenerateToken(user_found)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
