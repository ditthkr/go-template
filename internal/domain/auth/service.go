package auth

type Service interface {
	Register(username, email, password string) error
	Login(username, password string) (token string, err error)
}
