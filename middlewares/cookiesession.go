package middlewares

import (
	"time"

	"github.com/gobardofw/cache"
	"github.com/gobardofw/http/session"
	"github.com/gofiber/fiber/v2"
)

// NewCookieSession create new cookie based session
func NewCookieSession(cache cache.Cache, ctx *fiber.Ctx, secure bool, domain string, sameSite string, exp time.Duration) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s := session.NewCookieSession(cache, ctx, secure, domain, sameSite, exp, session.UUIDGenerator, "session")
		defer s.Save()
		s.Parse()
		ctx.Locals("session", s)
		return ctx.Next()
	}
}
