package user

import (
	"msg-app/src/api/models"
	"msg-app/src/api/repository/user"
)

type UserService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) CreateUser(inUser models.InUser) (models.OutUser, error) {
	return us.repository.CreateUser(inUser)
}

func (us *UserService) GetUsers() []models.OutUser {
	return us.repository.GetUsers()
}

// TODO
func (us *UserService) GetUserById() {
}

// TODO
func (us *UserService) UpdateUserById() {
}

// TODO
func (us *UserService) DeleteUser() {
}
