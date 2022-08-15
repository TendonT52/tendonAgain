package main

import (
	"log"
	"os"

	"github.com/TendonT52/tendon-api/configs"
	"github.com/TendonT52/tendon-api/drivers"
	"github.com/TendonT52/tendon-api/handlers"
	"github.com/TendonT52/tendon-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)




func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	App := configs.App{
		MongoDB: drivers.NewMongoDB(os.Getenv("DB_NAME"), os.Getenv("MONGO_DNS")),
		JwtSecret: services.NewJwtServices(os.Getenv("JWT_SECRET")),
	}	
	App.MongoDB.UserCollection = drivers.NewUserCollection(os.Getenv("USER_COLLECTION_NAME"), App.MongoDB)
	handlers.InitHadler(&App)	
	defer App.MongoDB.DisconnectMongo()

	r := gin.Default()
	r.POST("/signup", handlers.HandleSignUp)
	r.Run()

}
