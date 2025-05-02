package http

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"go-template/internal/adapter/http/auth"
	"go-template/internal/adapter/http/user"
	"go.uber.org/fx"
)

func NewApp() *fiber.App {
	return fiber.New()
}

func RegisterHTTPLifecycle(life fx.Lifecycle, app *fiber.App) {
	life.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				_ = app.Listen(":8080")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

var Module = fx.Options(
	fx.Provide(NewApp),
	fx.Invoke(RegisterHTTPLifecycle),
	fx.Invoke(
		auth.RegisterRoutes,
		user.RegisterRoutes,
	),
)
