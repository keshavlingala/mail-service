package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/keshavlingala/go-oauth2/config"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	println("Code: ", code)

	googlecon := config.GoogleConfig()

	// Get token to read, write and send email on behalf of user
	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	println("Token: ", token.AccessToken)
	println("Refresh Token: ", token.RefreshToken)
	println("Expiry: ", token.Expiry.String())

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	// Unmarshal the user data JSON string into a map
	var userData map[string]interface{}
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return c.SendString("User Data Unmarshal Failed")
	}

	response := make(map[string]interface{})
	response["token"] = token
	response["refreshToken"] = token.RefreshToken
	response["userData"] = userData
	return c.JSON(response)

}
