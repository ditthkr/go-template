package repository_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-template/internal/adapter/repository"
	"go-template/internal/domain/user"
	"go-template/internal/shared/testutil"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepository_Save(t *testing.T) {
	// Setup test database
	db := testutil.NewTestDB()
	repo := repository.NewUserRepository(db)

	t.Run("Save user successfully", func(t *testing.T) {
		// Create test user
		u := &user.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}

		// Save user
		err := repo.Save(u)

		// Assert
		require.NoError(t, err)
		assert.NotZero(t, u.Id, "User ID should be set after save")

		// Verify user exists in database
		var count int64
		db.Model(&struct{ ID uint64 }{}).Table("users").Where("id = ?", u.Id).Count(&count)
		assert.Equal(t, int64(1), count, "User should exist in database")
	})
}

func TestUserRepository_FindById(t *testing.T) {
	// Setup test database
	db := testutil.NewTestDB()
	repo := repository.NewUserRepository(db)

	// Create test user directly in DB
	testUser := testutil.CreateTestUser(t, db, "findbyid", "findbyid@example.com")

	t.Run("Find existing user by ID", func(t *testing.T) {
		// Find user
		found, err := repo.FindById(testUser.Id)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, testUser.Id, found.Id)
		assert.Equal(t, "findbyid", found.Username)
		assert.Equal(t, "findbyid@example.com", found.Email)
	})

	t.Run("Find non-existing user by ID", func(t *testing.T) {
		// Find non-existent user
		found, err := repo.FindById(999999)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, found)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestUserRepository_FindByUsername(t *testing.T) {
	// Setup test database
	db := testutil.NewTestDB()
	repo := repository.NewUserRepository(db)

	// Create test user directly in DB
	testutil.CreateTestUser(t, db, "findbyusername", "findbyusername@example.com")

	t.Run("Find existing user by username", func(t *testing.T) {
		// Find user
		found, err := repo.FindByUsername("findbyusername")

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, "findbyusername", found.Username)
		assert.Equal(t, "findbyusername@example.com", found.Email)
	})

	t.Run("Find non-existing user by username", func(t *testing.T) {
		// Find non-existent user
		found, err := repo.FindByUsername("nonexistent")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, found)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestUserRepository_ExistsUsername(t *testing.T) {
	// Setup test database
	db := testutil.NewTestDB()
	repo := repository.NewUserRepository(db)

	// Create test user directly in DB
	testutil.CreateTestUser(t, db, "existsusername", "existsusername@example.com")

	t.Run("Check existing username", func(t *testing.T) {
		// Check if username exists
		exists, err := repo.ExistsUsername("existsusername")

		// Assert
		require.NoError(t, err)
		assert.True(t, exists, "Username should exist")
	})

	t.Run("Check non-existing username", func(t *testing.T) {
		// Check if username exists
		exists, err := repo.ExistsUsername("nonexistent")

		// Assert
		require.NoError(t, err)
		assert.False(t, exists, "Username should not exist")
	})
}

func TestUserRepository_ExistsEmail(t *testing.T) {
	// Setup test database
	db := testutil.NewTestDB()
	repo := repository.NewUserRepository(db)

	// Create test user directly in DB
	testutil.CreateTestUser(t, db, "existsemail", "existsemail@example.com")

	t.Run("Check existing email", func(t *testing.T) {
		// Check if email exists
		exists, err := repo.ExistsEmail("existsemail@example.com")

		// Assert
		require.NoError(t, err)
		assert.True(t, exists, "Email should exist")
	})

	t.Run("Check non-existing email", func(t *testing.T) {
		// Check if email exists
		exists, err := repo.ExistsEmail("nonexistent@example.com")

		// Assert
		require.NoError(t, err)
		assert.False(t, exists, "Email should not exist")
	})
}
