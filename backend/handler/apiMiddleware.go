package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
	"time"
)

//ApiMiddleware is a middleware that checks if the user is authenticated and deserialize the token
func ApiMiddleware(ctx *fiber.Ctx) error {
	//Get the token from the cookie
	serialized_token := ctx.Cookies("token")

	//if token is blank, the user need to authenticate
	if serialized_token == "" {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "invalid cookie",
		})
	}

	//deserialize the token
	token, err := api.TokenDeserializer(serialized_token)

	//if there is an error or the token is nul, the user need to authenticate
	if err != nil || token == nil || token.AccessToken == "" {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "invalid cookie",
		})
	}

	//check if the token is expired
	if token.Expiry.Before(time.Now()) {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "token expired, please login again",
		})
	}

	//set the token in the context
	ctx.Locals("token", token)

	return ctx.Next()
}
