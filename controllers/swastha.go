package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keshavlingala/mail-service/models"
	"github.com/keshavlingala/mail-service/service"
)

func SwasthaContactUs(c *gin.Context) {

	var requestBody models.SwasthaBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Printf("Email: %s\nPhone: %s\nName: %s\nDescription: %s\n", requestBody.Email, requestBody.Phone, requestBody.Name, requestBody.Description)
	from := "Mail Assistant <davemicheaels@gmail.com>"
	to := "Swasthomeo@gmail.com"
	subject := "Email from SwastHomeo Website"
	body := fmt.Sprintf("Name: %s\r\nEmail: %s\nPhone Number: %s\nDescription: %s\n\n", requestBody.Name, requestBody.Email, requestBody.Phone, requestBody.Description)
	msg, err := service.SendMail(from, to, subject, body, requestBody.Email)
	if err != nil {
		c.String(500, "Error sending email")
		return
	}
	c.String(200, msg)
	return
}
