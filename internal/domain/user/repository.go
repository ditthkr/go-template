package user

type Repository interface {
	Save(*User) error
	FindById(id uint64) (*User, error)
	FindByUsername(username string) (*User, error)
	ExistsUsername(username string) (bool, error)
	ExistsEmail(email string) (bool, error)
}
