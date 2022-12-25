package repository

import (
	"strconv"
	"testing"

	m "msg-app/src/core/users/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	// Creating SQLite DB to test
	db, err := gorm.Open(sqlite.Open("../../../../database/mock.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("[test] failed to connect to sqlite")
	}
	// Creating a transaction to use in tests
	tx := db.Begin()
	tx.AutoMigrate(&m.User{})

	// Setting start point to rollback to
	tx.SavePoint("start")

	repository := NewUserRepository(tx)

	mockUser := m.RegisterUser{
		FirstName: "mockFirstName",
		LastName:  "mockLastName",
		Email:     "mockEmail",
		Password:  "mockPassword",
	}

	t.Run("Should register user", func(t *testing.T) {
		if err := repository.Register(mockUser); err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}
	})

	t.Run("Should get user by email", func(t *testing.T) {
		_, err := repository.GetUserByEmail(mockUser.Email)
		if err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}
	})

	t.Run("Should delete user by id", func(t *testing.T) {
		user, _ := repository.GetUserByEmail(mockUser.Email)
		id := strconv.FormatInt(user.ID, 10)
		if err := repository.DeleteUserById(id); err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}
		userDeleted, _ := repository.GetUserByEmail(mockUser.Email)
		if userDeleted.Email == mockUser.Email {
			t.Errorf("Have user = %v, Wanted user = nil", userDeleted)

		}

	})

	// Rollbacking to start and undoing all changes ocurred in tests
	tx.RollbackTo("start")
	tx.Commit()
}
