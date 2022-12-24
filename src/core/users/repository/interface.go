package repository

import m "msg-app/src/core/users/model"

type IUserRepository interface {
	Register(m.RegisterUser) error
	GetUserByEmail(string) (*m.User, error)
	UpdateUserById(string) error
	DeleteUserById(string) error
}
