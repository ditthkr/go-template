package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"go-template/internal/adapter/http/resp"
	"go-template/internal/domain/auth"
)

func RegisterRoutes(app *fiber.App, svc auth.Service, v *validator.Validate) {

	g := app.Group("/auth")

	g.Post("/signup", func(c fiber.Ctx) error {
		var req signUpReq
		if err := c.Bind().Body(&req); err != nil {
			return resp.Error(c, err)
		}
		if err := v.Struct(&req); err != nil {
			return resp.Error(c, err)
		}
		if err := svc.Register(req.Username, req.Email, req.Password); err != nil {
			return resp.Error(c, err)
		}
		return resp.Success(c, nil)
	})

	g.Post("/signin", func(c fiber.Ctx) error {
		var req signInReq
		if err := c.Bind().Body(&req); err != nil {
			return resp.Error(c, err)
		}
		if err := v.Struct(&req); err != nil {
			return resp.Error(c, err)
		}
		tok, err := svc.Login(req.Username, req.Password)
		if err != nil {
			return resp.Error(c, err)
		}
		return resp.Success(c, map[string]string{"token": tok})
	})
}
