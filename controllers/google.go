package controllers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/keshavlingala/mail-service/config"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

func GoogleLogin(c *gin.Context) {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)

	c.Status(fiber.StatusSeeOther)
	// Redirect the user to the Google login page
	c.Redirect(http.StatusSeeOther, url)
	return
}

func GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		return
	}

	code := c.Query("code")

	println("Code: ", code)

	googlecon := config.GoogleConfig()

	// Get token to read, write and send email on behalf of user
	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	println("Token: ", token.AccessToken)
	println("Refresh Token: ", token.RefreshToken)
	println("Expiry: ", token.Expiry.String())

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal the user data JSON string into a map
	var userData map[string]interface{}
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return
	}

	response := make(map[string]interface{})
	response["token"] = token
	response["refreshToken"] = token.RefreshToken
	response["userData"] = userData
	c.JSON(http.StatusOK, response)
	return

}
