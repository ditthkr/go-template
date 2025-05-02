package resp

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"go-template/internal/shared/errs"
)

type successResp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type errorResp struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func Error(c fiber.Ctx, err error) error {
	var code int
	switch {
	case errors.Is(err, errs.ErrInvalidCredential):
		code = fiber.StatusUnauthorized // 401
	case errors.Is(err, errs.ErrDuplicateUsername):
		code = fiber.StatusConflict // 409
	case errors.Is(err, errs.ErrDuplicateEmail):
		code = fiber.StatusConflict // 409
	default:
		code = fiber.StatusInternalServerError // 500
	}

	return c.Status(code).JSON(errorResp{
		Success: false,
		Message: err.Error(),
	})
}

func Success(c fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(successResp{
		Success: true,
		Data:    data,
	})
}
