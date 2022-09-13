package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
)

// GoogleAuth is a handler for the /api/auth route.
func GoogleAuth(c *fiber.Ctx) error {
	// redirect to the Google OAuth2 consent page
	return c.Redirect(api.GetAuthUrl())
}

// GoogleCallback is a handler for the /api/auth/callback route.
func GoogleCallback(c *fiber.Ctx) error {
	// get the code from the query string
	code := c.Query("code")

	// exchange the code for a token
	_, err := api.CodeToToken(code)
	if err != nil {
		return err
	}

	// redirect to the home page
	return c.Redirect("/")
}
