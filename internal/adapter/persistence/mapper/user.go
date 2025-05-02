package mapper

import (
	"go-template/internal/adapter/persistence/model"
	"go-template/internal/domain/user"
)

func UserToModel(r *user.User) *model.User {
	return &model.User{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
	}
}

func UserToDomain(r *model.User) *user.User {
	return &user.User{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Status:   r.Status,
	}
}
