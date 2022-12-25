package service

import m "msg-app/src/core/users/model"

type IUserService interface {
	Register(m.RegisterUser) error
	Login(m.LoginUser) bool
	GetUserByEmail(string) (*m.OutUser, error)
	// UpdateUserById(string) error
	// DeleteUserById(string) error
}
