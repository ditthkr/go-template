package auth_test

import (
	"github.com/golang/mock/gomock"
	"go-template/internal/shared"
	"go-template/internal/shared/errs"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-template/internal/service/auth"
	"go-template/internal/service/user/mock"
)

func TestAuthService_Register(t *testing.T) {
	t.Run("Register success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock.NewMockRepository(ctrl)
		sessionStore := mock.NewMockStore(ctrl)
		jwt := new(shared.JWT)
		service := auth.NewService(userRepo, sessionStore, jwt)

		userRepo.EXPECT().ExistsUsername("testuser").Return(false, nil)
		userRepo.EXPECT().ExistsEmail("test@example.com").Return(false, nil)
		userRepo.EXPECT().Save(gomock.Any()).Return(nil)

		err := service.Register("testuser", "test@example.com", "Password123")
		assert.NoError(t, err)
	})

	t.Run("Username already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock.NewMockRepository(ctrl)
		sessionStore := mock.NewMockStore(ctrl)
		jwt := new(shared.JWT)
		service := auth.NewService(userRepo, sessionStore, jwt)

		userRepo.EXPECT().ExistsUsername("testuser").Return(true, nil)

		err := service.Register("testuser", "test@example.com", "Password123")

		assert.Error(t, err)
		assert.Equal(t, errs.ErrDuplicateUsername, err)
		assert.Contains(t, err.Error(), "username already exists")
	})

	t.Run("Email already exists", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock.NewMockRepository(ctrl)
		sessionStore := mock.NewMockStore(ctrl)
		jwt := new(shared.JWT)
		service := auth.NewService(userRepo, sessionStore, jwt)

		userRepo.EXPECT().ExistsUsername("testuser").Return(false, nil)
		userRepo.EXPECT().ExistsEmail("test@example.com").Return(true, nil)

		err := service.Register("testuser", "test@example.com", "Password123")

		assert.Error(t, err)
		assert.Equal(t, errs.ErrDuplicateEmail, err)
		assert.Contains(t, err.Error(), "email already exists")
	})
}
