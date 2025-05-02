package app

import (
	"go-template/internal/adapter/http"
	"go-template/internal/adapter/repository"
	"go-template/internal/service"
	"go-template/internal/shared"
	"go.uber.org/fx"
)

var Module = fx.Options(
	shared.Module,
	repository.Module,
	service.Module,
	http.Module,
)
