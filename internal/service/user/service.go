package user

import (
	"go-template/internal/domain/user"
)

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) user.Service {
	return &service{repo: repo}
}

func (r *service) Profile(id uint64) (*user.User, error) {
	return r.repo.FindById(id)
}
