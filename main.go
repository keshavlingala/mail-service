package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/keshavlingala/go-oauth2/config"
	"github.com/keshavlingala/go-oauth2/controllers"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"os"
)

type KeshavBody struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Content  string      `json:"content"`
	UserData interface{} `json:"userData"`
}

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/keshav", func(c *fiber.Ctx) error {

		// Get the user's email and content from the request.
		var request KeshavBody
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request.")
		}

		fmt.Println("Email: ", request.Email)
		fmt.Println("Content: ", request.Content)
		fmt.Println("Name: ", request.Name)
		fmt.Println("User Data: ", request.UserData)

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
		body := fmt.Sprintf(`
	Name: %s
	Email: %s
	Message: %s
	Aditional User Data:
%s

	Full Request: 
%s
`, request.Name, request.Email, request.Content, ToJson(request.UserData), c.Request().String())
		message.Raw = base64.URLEncoding.EncodeToString([]byte(
			fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body),
		))

		// Send the email.
		_, err = mailService.Users.Messages.Send("me", &message).Do()
		if err != nil {
			log.Fatalf("Unable to send email: %v", err)
		}

		fmt.Println("Email sent successfully!")
		return c.SendString("Email sent successfully!")
	})

	app.Listen(":8080")

}

func ToJson(data interface{}) string {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return ""
	}
	return string(json)
}
