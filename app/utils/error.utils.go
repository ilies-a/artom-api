package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, code int, msg string) bool {
	if err == nil {
		return false
	}
	c.JSON(code, msg)
	return true
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func DefaultErrorMsg(code int) string {
	switch code {
	case 400:
		return "bad request"
	case 404:
		return "not found"
	case 500:
		return "server error"
	default:
		return "unkown error"
	}
}
