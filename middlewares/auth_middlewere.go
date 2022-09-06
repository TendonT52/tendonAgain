package middlewares

import (
	"net/http"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	app *configs.App
}

func NewMiddlewares(app *configs.App) *Middlewares {
	return &Middlewares{
		app: app,
	}
}

func (middlewere Middlewares) AuthMiddlewere(c *gin.Context) {

	token, err := c.Cookie("token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	id, err := middlewere.app.JwtSecret.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	c.Set("userId", id)

	c.Next()
}
