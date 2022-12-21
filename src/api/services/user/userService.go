package user

import (
	m "msg-app/src/api/models"
	"msg-app/src/api/repository/user"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) CreateUser(inUser m.InUser) (m.OutUser, error) {
	return us.repository.CreateUser(inUser)
}

func (us *UserService) GetUserById(id string) m.OutUser {
	return us.repository.GetUserById(id)
}

// TODO
func (us *UserService) UpdateUserById() {
}

// TODO
func (us *UserService) DeleteUser() {
}
