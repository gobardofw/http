package middlewares

import (
	"github.com/gobardofw/http/session"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CSRFMiddleware protection middleware
func CSRFMiddleware(ctx *fiber.Ctx) error {
	session, ok := ctx.Locals("session").(session.Session)
	if !ok {
		panic("CSRF: Failed access csrf token on session!")
	}

	token, _ := session.Get("csrf_token").(string)
	if token == "" {
		session.Set("csrf_token", uuid.New().String())
	}

	return ctx.Next()
}

// CSRFMiddleware
func GetCSRFKey(ctx *fiber.Ctx) string {
	session, ok := ctx.Locals("session").(session.Session)
	if ok {
		if token, ok := session.Get("csrf_token").(string); ok {
			return token
		}
	}
	return ""
}
