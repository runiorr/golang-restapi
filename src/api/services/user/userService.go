package user

import (
	"fmt"
	m "msg-app/src/api/models"
	"msg-app/src/api/repository/user"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) CreateUser(inUser m.InUser) error {
	if err := us.repository.CreateUser(inUser); err != nil {
		return &InternalError{err}
	}
	return nil
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
