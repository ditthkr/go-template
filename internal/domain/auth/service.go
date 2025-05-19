package auth

import "context"

type Service interface {
	Register(ctx context.Context, username, email, password string) error
	Login(ctx context.Context, username, password string) (token string, err error)
}
