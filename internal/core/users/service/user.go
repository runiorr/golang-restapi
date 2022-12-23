package service

import (
	"fmt"
	m "msg-app/internal/core/users/model"
	user "msg-app/internal/core/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) Register(registerUser m.RegisterUser) error {
	if err := us.repository.Register(registerUser); err != nil {
		return &InternalError{err}
	}
	return nil
}

func (us *UserService) Login(loginUser m.LoginUser) bool {
	dbUser, err := us.repository.GetUserByEmail(loginUser.Email)
	if err != nil {
		return false
	}

	userPass := []byte(loginUser.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	return passErr == nil
}

func (us *UserService) GetUserByEmail(email string) (*m.OutUser, error) {
	user, err := us.repository.GetUserByEmail(email)
	if err != nil {
		return nil, &InternalError{err}
	}

	outUser := m.OutUser{
		Id:        fmt.Sprint(user.ID),
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	return &outUser, nil
}

// Todo
func (us *UserService) UpdateUserById(id string) string {
	return us.repository.UpdateUserById(id)
}

// Todo
func (us *UserService) DeleteUserById(id string) string {
	return us.repository.DeleteUserById(id)
}
