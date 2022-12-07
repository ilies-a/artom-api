package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilies-a/artom-api/app"
	"github.com/ilies-a/artom-api/app/services"
	"github.com/joho/godotenv"
)

func main() {
	configEnv()
	services.ConnectDB()
	defer services.CloseDB()
	app.InitServer()
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func configEnv() {
	switch os.Getenv("APP_ENV") {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		loadDotEnv()
	}
}
