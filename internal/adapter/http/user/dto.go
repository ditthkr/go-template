package user

import "time"

type userResp struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
