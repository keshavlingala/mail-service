package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keshavlingala/mail-service/config"
	"github.com/keshavlingala/mail-service/models"
	"github.com/keshavlingala/mail-service/utils"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func KeshavEmail(c *gin.Context) {

	// Get the user's email and content from the request.

	var request models.PortfolioBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Println("Email: ", request.Email)
	log.Println("Content: ", request.Content)
	log.Println("Name: ", request.Name)
	log.Println("User Data: ", request.UserData)

	config := config.GoogleConfig()

	token, err := config.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN"),
	}).Token()

	client := config.Client(context.Background(), token)

	mailService, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		print("Error creating mail service")
	}

	// Compose the email.
	var message gmail.Message
	from := "Mail Assistant <davemicheaels@gmail.com>"
	to := "keshavlingala@gmail.com"
	subject := "Email from keshav.codes"
	requestDetails := fmt.Sprintf("UserAgent: %s\nIP: %s\nHost: %s\n", c.GetHeader("User-Agent"), c.ClientIP(), c.Request.Host)

	body := fmt.Sprintf("Name: %s\r\nEmail: %s\nMessage: %s\nAditional User Data:%s\n\nRequest Details: %s", request.Name, request.Email, request.Content, utils.ToJson(request.UserData), requestDetails)
	message.Raw = base64.URLEncoding.EncodeToString([]byte(
		fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body),
	))

	// Send the email.
	_, err = mailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	c.String(http.StatusOK, "Email sent successfully!")
	return
}
