package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/handlers"
	"github.com/TendonT52/tendon-api/middlewares"
	"github.com/TendonT52/tendon-api/models"
	"github.com/TendonT52/tendon-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app configs.App

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongo, err := models.NewMongoDB(
		os.Getenv("DB_NAME"),
		os.Getenv("USER_COLLECTION_NAME"),
		os.Getenv("JWT_COLLECTION_NAME"),
		os.Getenv("MONGO_DNS"),
	)
	if err != nil {
		log.Fatal(err)
	}
	jwt := services.NewJwtServices(os.Getenv("JWT_ACCESS_SECRET"))
	app = configs.App{
		MongoDB:   mongo,
		JwtSecret: jwt,
	}
}

func main() {
	app = *configs.NewApp()
	authHandler := handlers.NewHandlerAuth(&app)
	curriculumHandler := handlers.NewCurriculumHandler(&app)
	middlewares := middlewares.NewMiddlewares(&app)

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "ok1"}) })
	r.POST("/signup", authHandler.HandleSignUp)
	r.POST("/signin", authHandler.HandleSignIn)
	r.GET("/check")
	rg1 := r.Group("/auth", middlewares.AuthMiddlewere)
	rg1.GET("/curriculum", curriculumHandler.GetCurriculum)
	r.Run()
}
