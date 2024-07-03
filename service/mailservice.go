package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/keshavlingala/mail-service/config"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
	"os"
)

func SendMail(from string, to string, subject string, body string, replyTo string) (string, error) {

	gconfig := config.GoogleConfig()

	token, err := gconfig.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN"),
	}).Token()

	if err != nil {
		return "", err
	}

	client := gconfig.Client(context.Background(), token)

	mailService, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		print("Error creating mail service")
		return "", err
	}

	// Compose the email.
	var message gmail.Message
	message.Raw = base64.URLEncoding.EncodeToString([]byte(
		fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body),
	))

	if replyTo != "" {
		message.Raw = base64.URLEncoding.EncodeToString([]byte(
			fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nReply-To: %s\r\n\r\n%s", from, to, subject, replyTo, body),
		))
	}

	// Send the email.
	_, err = mailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
		return "", err
	}
	return "Email sent successfully", nil
}
