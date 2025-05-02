package session

type Store interface {
	Set(uid string, jti string) error
	Get(uid string) (jti string, ok bool)
}
