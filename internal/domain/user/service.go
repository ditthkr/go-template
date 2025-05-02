package user

type Service interface {
	Profile(id uint64) (*User, error)
}
