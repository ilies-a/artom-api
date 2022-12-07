package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilies-a/artom-api/app/models"
	"github.com/ilies-a/artom-api/app/services"
	. "github.com/ilies-a/artom-api/app/utils"
)

func GetAllEmails(c *gin.Context) {
	emails, err := models.GetAllEmails()
	if err != nil {
		c.JSON(500, DefaultErrorMsg(500))
		return
	}
	c.JSON(200, emails)
}

func SaveEmail(c *gin.Context) {
	var email models.EmailRequestStructure

	err := c.ShouldBindJSON(&email)
	if err != nil {
		c.JSON(400, DefaultErrorMsg(400))
		return
	}

	err = models.SaveEmail(email.Email)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"emails_value_key\"" {
			c.JSON(403, "Email already exists")
		} else {
			c.JSON(500, DefaultErrorMsg(500))
		}
		return
	}

	err = services.SendWelcomeEmail(email.Email)
	if err != nil {
		c.JSON(201, "Successfully added "+email.Email+" but failed sending him email")
	} else {
		c.JSON(201, "Successfully added "+email.Email+" and successfully sent him email")
	}
}

func DeleteEmail(c *gin.Context) {
	emailId := c.Param("id")

	err := models.DeleteEmail(emailId)
	if err != nil {
		c.JSON(500, DefaultErrorMsg(500))
		return
	}

	c.JSON(200, "Successfully deleted email")
}
