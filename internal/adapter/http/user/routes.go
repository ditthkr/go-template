package user

import (
	"github.com/gofiber/fiber/v3"
	"go-template/internal/adapter/http/middleware"
	"go-template/internal/adapter/http/resp"
	"go-template/internal/domain/session"
	"go-template/internal/domain/user"
	"go-template/internal/shared"
	"strconv"
)

func RegisterRoutes(app *fiber.App, svc user.Service, store session.Store, jwt *shared.JWT) {

	g := app.Group("/users", middleware.Auth(store, jwt))

	g.Get("/me", func(c fiber.Ctx) error {
		uidStr := c.Locals("uid").(string)
		uid, _ := strconv.ParseUint(uidStr, 10, 64)

		u, err := svc.Profile(uid)
		if err != nil {
			return resp.Error(c, err) // ใช้ helper JSON error
		}
		return resp.Success(c, userResp{
			Id:        u.Id,
			Username:  u.Username,
			Email:     u.Email,
			Status:    false,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	})
}
