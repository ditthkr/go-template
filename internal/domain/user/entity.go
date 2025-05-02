package user

import "time"

type User struct {
	Id        uint64
	Username  string
	Email     string
	Password  string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
