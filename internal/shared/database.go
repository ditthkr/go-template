package shared

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
}
