package service

import (
	"errors"
	m "msg-app/src/core/users/model"
	"strings"
	"testing"
)

type MockRepository struct {
	registerFn       func(registerUser m.RegisterUser) error
	getUserByEmailFn func(email string) (*m.User, error)
	// updateUserByIdFn func(id string) error
	// deleteUserByIdFn func(id string) error
}

func (r *MockRepository) Register(registerUser m.RegisterUser) error {
	return r.registerFn(registerUser)
}

func (r *MockRepository) GetUserByEmail(email string) (*m.User, error) {
	return r.getUserByEmailFn(email)
}

// func (r *MockRepository) UpdateUserById(id string) error {
// 	return r.updateUserByIdFn(id)
// }

// func (r *MockRepository) DeleteUserById(id string) error {
// 	return r.deleteUserByIdFn(id)
// }

func TestUserService(t *testing.T) {
	mockedUsers := map[int]*m.User{
		1: {ID: 1,
			FirstName: "",
			LastName:  "",
			Email:     "",
			Password:  "",
		}}

	// Mocking repository methods
	mockRepository := &MockRepository{
		registerFn: func(registerUser m.RegisterUser) error {
			if registerUser.FirstName == "" ||
				registerUser.LastName == "" ||
				registerUser.Email == "" ||
				registerUser.Password == "" {
				return errors.New("Fill all fields.")
			}

			for _, user := range mockedUsers {
				if user.Email == registerUser.Email {
					return errors.New("User email is in use.")
				}
			}

			user := m.User{
				ID:        2,
				FirstName: registerUser.FirstName,
				LastName:  registerUser.LastName,
				Email:     registerUser.Email,
				Password:  registerUser.Password,
			}

			mockedUsers[2] = &user

			return nil
		},
		getUserByEmailFn: func(email string) (*m.User, error) {
			for _, user := range mockedUsers {
				if user.Email == email {
					return user, nil
				}
			}
			return nil, errors.New("User not found")
		},
	}

	service := NewUserService(mockRepository)

	mockedUser := m.RegisterUser{
		FirstName: "name",
		LastName:  "surname",
		Email:     "email@test.com",
		Password:  "test123",
	}

	t.Run("Should register user", func(t *testing.T) {
		if err := service.Register(mockedUser); err != nil {
			t.Errorf("Have error = %v, Wanted error = nil", err)
		}
	})

	t.Run("Should return internal error", func(t *testing.T) {
		user := mockedUser
		user.Email = ""

		if err := service.Register(user); err != nil {
			splitError := strings.Split(err.Error(), ":")
			if !strings.Contains(splitError[0], "Status") {
				t.Errorf("Have error = %v, Wanted error to contain status", splitError[0])
			}
		}
	})

	t.Run("Should find user by email", func(t *testing.T) {
		_, err := service.GetUserByEmail(mockedUser.Email)
		if err != nil {
			t.Errorf("Have error = %v, Wanted err = nil", err)
		}
	})

	t.Run("Should return error in find user by email", func(t *testing.T) {
		_, err := service.GetUserByEmail("fake@email")
		if err == nil {
			t.Errorf("Have error = %v, Wanted err = internal", err)
		}
	})

	t.Run("Should login user", func(t *testing.T) {
		loginUser := m.LoginUser{
			Email:    mockedUser.Email,
			Password: mockedUser.Password,
		}

		ok := service.Login(loginUser)
		if !ok {
			t.Errorf("Have ok = %v, Wanted ok = true", ok)
		}
	})

	t.Run("Should return wrong login (email)", func(t *testing.T) {
		loginUser := m.LoginUser{
			Email:    "WrongEmail",
			Password: mockedUser.Password,
		}

		ok := service.Login(loginUser)
		if ok {
			t.Errorf("Have ok = %v, Wanted ok = false", ok)
		}
	})

	t.Run("Should return wrong login (password)", func(t *testing.T) {
		loginUser := m.LoginUser{
			Email:    mockedUser.Email,
			Password: "WrongPassword",
		}

		ok := service.Login(loginUser)
		if ok {
			t.Errorf("Have ok = %v, Wanted ok = false", ok)
		}
	})

}
