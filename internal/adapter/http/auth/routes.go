package auth

import (
	"github.com/ditthkr/loggie"
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

		ctx := c.Context()
		log := loggie.FromContext(ctx)
		log.Info("received /signup request")

		if err := svc.Register(ctx, req.Username, req.Email, req.Password); err != nil {
			return resp.Error(c, err)
		}
		return resp.Success(c, nil)
	})

	g.Post("/signin", func(c fiber.Ctx) error {

		ctx := c.Context()
		ctx = loggie.WithCustomField(ctx, "user_id", c.Locals("uid"))
		log := loggie.FromContext(ctx)
		log.Info("received /signin request")

		var req signInReq
		if err := c.Bind().Body(&req); err != nil {
			return resp.Error(c, err)
		}
		if err := v.Struct(&req); err != nil {
			return resp.Error(c, err)
		}
		tok, err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return resp.Error(c, err)
		}
		return resp.Success(c, map[string]string{"token": tok})
	})
}
