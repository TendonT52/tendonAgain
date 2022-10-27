package handlers

import (
	"net/http"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/gin-gonic/gin"
)

type CurriculumHandler struct {
	app *configs.App
}

func NewCurriculumHandler(app *configs.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (handleAuth *AuthHandler) GetCurriculum(ctx *gin.Context) {
	userId, ok := ctx.Get("userId");
	if !ok{
		ctx.JSON(http.StatusBadRequest, userId)
	}
	ctx.JSON(http.StatusAccepted, userId)
}