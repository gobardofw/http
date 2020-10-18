package middlewares

import (
	"time"

	"github.com/gobardofw/cache"
	"github.com/gobardofw/http/session"
	"github.com/gofiber/fiber/v2"
)

// NewHeaderSession create new header based session
func NewHeaderSession(cache cache.Cache, exp time.Duration) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s := session.NewHeaderSession(cache, ctx, exp, session.UUIDGenerator, "X-SESSION-ID")
		defer s.Save()
		s.Parse()
		ctx.Locals("session", s)
		return ctx.Next()
	}
}
