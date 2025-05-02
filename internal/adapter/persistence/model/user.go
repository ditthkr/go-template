package model

import "time"

type User struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	Status    bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
