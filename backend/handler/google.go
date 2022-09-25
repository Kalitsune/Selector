package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
)

// GoogleAuth is a handler for the /api/auth route.
func GoogleAuth(ctx *fiber.Ctx) error {
	// redirect to the Google OAuth2 consent page
	return ctx.Redirect(api.GetAuthUrl(), 308)
}

// GoogleCallback is a handler for the /api/auth/callback route.
func GoogleCallback(ctx *fiber.Ctx) error {
	// get the code from the query string
	code := ctx.Query("code")

	// exchange the code for a token
	token, err := api.CodeToToken(code)
	if err != nil {
		return err
	}

	//serialize the token
	serializedToken, err := api.TokenSerializer(token)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	//store the cookie
	cookie := &fiber.Cookie{
		Name:    "token",
		Value:   serializedToken,
		Expires: token.Expiry,
	}
	ctx.Cookie(cookie)

	// load the callback page
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return ctx.SendString("<p>Success! You'll soon be redirected.</p><script>window.opener.postMessage('success', '*')</script>")
}
