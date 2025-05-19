package http

import (
	"context"
	"github.com/ditthkr/loggie"
	"github.com/gofiber/fiber/v3"
	"go-template/internal/adapter/http/auth"
	"go-template/internal/adapter/http/user"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewApp() *fiber.App {
	return fiber.New()
}

func RegisterHTTPLifecycle(life fx.Lifecycle, app *fiber.App, logger *zap.Logger) {

	app.Use(func(c fiber.Ctx) error {
		ctx, traceId := loggie.Injection(c.Context(), &loggie.ZapLogger{L: logger})
		c.SetContext(ctx)
		c.Set("X-Trace-Id", traceId)
		return c.Next()
	})
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
	fx.Provide(NewApp, func() (*zap.Logger, error) {
		return zap.NewProduction(zap.AddCallerSkip(1))
	}),
	fx.Invoke(RegisterHTTPLifecycle),
	fx.Invoke(
		auth.RegisterRoutes,
		user.RegisterRoutes,
	),
)
