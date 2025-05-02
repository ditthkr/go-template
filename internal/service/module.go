package service

import (
	"go-template/internal/service/auth"
	"go-template/internal/service/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		auth.NewService,
		user.NewService,
	))
