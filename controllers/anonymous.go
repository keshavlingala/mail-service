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

func AnonymousEmail(c *gin.Context) {
	var request models.AnonymousBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Println("Content: ", request.Content)
	log.Println("Meta Data: ", request.MetaData)

	from := "Mail Assistant <davemicheaels@gmail.com>"
	to := "Keshav Lingala <admin@keshav.codes>"
	subject := "New Anonymous Email from keshav.codes"
	requestDetails := fmt.Sprintf("UserAgent: %s\nIP: %s\nHost: %s\n", c.GetHeader("User-Agent"), c.ClientIP(), c.Request.Host)
	body := fmt.Sprintf("Message: %s\nAditional Meta Data:%s\n\nRequest Details: %s", request.Content, utils.ToJson(request.MetaData), requestDetails)
	msg, err := service.SendMail(from, to, subject, body, "")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error sending email")
		return
	}

	fmt.Println(msg)
	c.String(http.StatusOK, msg)
	return
}
