package repository

import (
	"go-template/internal/adapter/persistence/mapper"
	"go-template/internal/adapter/persistence/model"
	"go-template/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Save(u *user.User) error {
	m := model.User{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	u.Id = m.Id
	return nil
}

func (r *UserRepository) FindById(id uint64) (*user.User, error) {
	var m model.User
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return mapper.UserToDomain(&m), nil
}

func (r *UserRepository) FindByUsername(username string) (*user.User, error) {
	var m model.User
	if err := r.db.First(&m, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return mapper.UserToDomain(&m), nil
}

func (r *UserRepository) ExistsUsername(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) ExistsEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var _ user.Repository = (*UserRepository)(nil)
