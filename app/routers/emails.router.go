package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilies-a/artom-api/app/handlers"
	"github.com/ilies-a/artom-api/app/middlewares"
)

func InitEmailsRouter(r *gin.Engine) {
	emailsR := r.Group("/emails")
	{
		emailsR.POST("/", handlers.SaveEmail)
		emailsR.Use(middlewares.BasicAuth())
		emailsR.GET("/", handlers.GetAllEmails)
		emailsR.DELETE("/:id", handlers.DeleteEmail)
	}
}
