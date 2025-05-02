package auth

import (
	"go-template/internal/domain/auth"
	"go-template/internal/domain/session"
	"go-template/internal/domain/user"
	"go-template/internal/shared"
	"go-template/internal/shared/errs"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type service struct {
	userRepo     user.Repository
	sessionStore session.Store
	jwt          *shared.JWT
}

func NewService(u user.Repository, st session.Store, j *shared.JWT) auth.Service {
	return &service{userRepo: u, sessionStore: st, jwt: j}
}

func (r *service) Register(username, email, password string) error {

	username = strings.TrimSpace(username)
	email = strings.ToLower(strings.TrimSpace(email))

	if ok, _ := r.userRepo.ExistsUsername(username); ok {
		return errs.ErrDuplicateUsername
	}
	if ok, _ := r.userRepo.ExistsEmail(email); ok {
		return errs.ErrDuplicateEmail
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return r.userRepo.Save(&user.User{
		Username: username, Email: email, Password: string(hash),
	})
}

func (r *service) Login(username, password string) (string, error) {
	u, err := r.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return "", errs.ErrInvalidCredential
	}
	token, jti, err := r.jwt.IssueToken(strconv.FormatUint(u.Id, 10))
	if err != nil {
		return "", err
	}
	err = r.sessionStore.Set(strconv.FormatUint(u.Id, 10), jti)
	if err != nil {
		return "", err
	}
	return token, nil
}
