package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilies-a/artom-api/app/middlewares"
	"github.com/ilies-a/artom-api/app/routers"
)

func InitServer() {

	r := gin.New()

	r.Use(gin.Recovery(), middlewares.CORSMiddleware())

	routers.InitEmailsRouter(r)

	port := os.Getenv("PORT")

	r.Run("localhost:" + port)
}
