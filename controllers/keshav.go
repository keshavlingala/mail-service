package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keshavlingala/mail-service/models"
	"github.com/keshavlingala/mail-service/service"
	"github.com/keshavlingala/mail-service/utils"
	"log"
	"net/http"
)

func KeshavEmail(c *gin.Context) {
	var request models.PortfolioBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Println("Email: ", request.Email)
	log.Println("Content: ", request.Content)
	log.Println("Name: ", request.Name)
	log.Println("User Data: ", request.UserData)

	from := "Mail Assistant <davemicheaels@gmail.com>"
	to := "keshavlingala@gmail.com"
	subject := "Email from keshav.codes"
	requestDetails := fmt.Sprintf("UserAgent: %s\nIP: %s\nHost: %s\n", c.GetHeader("User-Agent"), c.ClientIP(), c.Request.Host)
	body := fmt.Sprintf("Name: %s\r\nEmail: %s\nMessage: %s\nAditional User Data:%s\n\nRequest Details: %s", request.Name, request.Email, request.Content, utils.ToJson(request.UserData), requestDetails)
	msg, err := service.SendMail(from, to, subject, body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error sending email")
		return
	}

	fmt.Println(msg)
	c.String(http.StatusOK, msg)
	return
}
