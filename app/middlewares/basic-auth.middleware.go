package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("ADMIN_USERNAME"): os.Getenv("ADMIN_PASSWORD"),
	})
}
