package repository

import (
	"context"
	"go-template/internal/adapter/persistence/model"
	"go-template/internal/domain/session"
	"go-template/internal/domain/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func RegisterMigration(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.AutoMigrate(
				&model.User{},
			)
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		NewSession,
		fx.Annotate(NewSession, fx.As(new(session.Store))),
		NewUserRepository,
		fx.Annotate(NewUserRepository, fx.As(new(user.Repository))),
	),
	fx.Invoke(RegisterMigration),
)
