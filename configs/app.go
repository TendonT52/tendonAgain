package configs

import (
	"log"
	"os"

	"github.com/TendonT52/tendon-api/models"
	"github.com/TendonT52/tendon-api/services"
)

type App struct {
	MongoDB   *models.DB
	JwtSecret *services.JwtServices
}

func NewApp() *App {
	mongo, err := models.NewMongoDB(os.Getenv("DB_NAME"), os.Getenv("USER_COLLECTION_NAME"),os.Getenv("MONGO_DNS"))
	if err != nil {
		log.Fatal(err)
	}
	jwtServices := services.NewJwtServices(os.Getenv("JWT_ACCESS_SECRET"))
	return &App{
		MongoDB: mongo,
		JwtSecret: jwtServices,
	}
}