package repository

import (
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
		tx.SavePoint("register")

		if err := repository.Register(mockUser); err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}

		tx.RollbackTo("register")
	})

	t.Run("Should get user by email", func(t *testing.T) {
		tx.SavePoint("getUserByEmail")
		repository.Register(mockUser)

		_, err := repository.GetUserByEmail(mockUser.Email)
		if err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}
		tx.RollbackTo("getUserByEmail")
	})

	t.Run("Should delete user by id", func(t *testing.T) {
		tx.SavePoint("deleteUserById")
		repository.Register(mockUser)

		user, _ := repository.GetUserByEmail(mockUser.Email)
		if err := repository.DeleteUserById(user.ID); err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}

		tx.RollbackTo("deleteUserById")
	})

	// Rollbacking to start and undoing all changes ocurred in tests
	tx.RollbackTo("start")
	tx.Commit()
}
