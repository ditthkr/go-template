package testutil

import (
	"github.com/stretchr/testify/require"
	"go-template/internal/adapter/persistence/model"
	"go-template/internal/domain/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func NewTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}

	// Auto-migrate user model for testing
	// Note: This creates a simplified users table for testing
	err = db.AutoMigrate(&model.User{})

	if err != nil {
		panic("failed to migrate test database: " + err.Error())
	}

	return db
}

func CreateTestUser(t *testing.T, db *gorm.DB, username, email string) *user.User {
	// Create a user record
	userModel := &model.User{
		Username: username,
		Email:    email,
		Password: "hashedpassword",
	}

	err := db.Create(userModel).Error
	require.NoError(t, err)

	// Convert to domain user
	return &user.User{
		Id:       userModel.Id,
		Username: userModel.Username,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}
