package shared

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(LoadConfig, NewDatabase, NewRedis, NewJWT, NewValidator),
)
