package repository

import m "msg-app/src/core/users/model"

type IUserRepository interface {
	Register(registerUser m.RegisterUser) error
	GetUserByEmail(email string) (*m.User, error)
	// UpdateUserById(id string) error
	// DeleteUserById(d string) error
}
