package middleware

import (
	"github.com/gofiber/fiber/v3"
	"go-template/internal/domain/session"
	"go-template/internal/shared"
	"strings"
)

func Auth(store session.Store, jwt *shared.JWT) fiber.Handler {
	return func(c fiber.Ctx) error {
		h := c.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			return fiber.ErrUnauthorized
		}
		claims, err := jwt.ParseToken(h[7:])
		if err != nil {
			return fiber.ErrUnauthorized
		}

		if j, ok := store.Get(claims.UserID); !ok || j != claims.ID {
			return fiber.ErrUnauthorized // token เก่าถูกแทน
		}
		c.Locals("uid", claims.UserID)
		return c.Next()
	}
}
