package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
	"time"
)

// ApiMiddleware is a middleware that checks if the user is authenticated and deserialize the token
func ApiMiddleware(ctx *fiber.Ctx) error {
	//Get the token from the cookie
	serializedToken := ctx.Cookies("token")

	//if token is blank, the user need to authenticate
	if serializedToken == "" {
		api.Logger.Warning.Println("a request has been canceled: No cookie")
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "no cookie",
		})
	}

	//deserialize the token
	token, err := api.TokenDeserializer(serializedToken)

	//if there is an error or the token is nul, the user need to authenticate
	if err != nil || token == nil || token.AccessToken == "" {
		api.Logger.Warning.Println("a request has been canceled: invalid cookie")

		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"error": "invalid cookie",
		})
	}

	//check if the token is expired
	if token.Expiry.Before(time.Now()) {

		//refresh the token
		token, err = api.RefreshToken(token)

		//if there is an error, the user need to authenticate again
		if err != nil {
			api.Logger.Warning.Printf("a request has been canceled: token is expired: %v", err)

			ctx.Status(fiber.StatusUnauthorized)
			return ctx.JSON(fiber.Map{
				"error": "token expired, please login again",
			})
		}
	}

	//set the token in the context
	ctx.Locals("token", token)

	return ctx.Next()
}
