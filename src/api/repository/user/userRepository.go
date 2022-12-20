package user

import (
	"msg-app/src/api/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	data map[string]models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{data: make(map[string]models.User)}
}

func (ur *UserRepository) CreateUser(inUser models.InUser) (models.OutUser, error) {
	id := uuid.New().String()

	user := models.User{
		Id:        id,
		FirstName: inUser.FirstName,
		LastName:  inUser.LastName,
		Email:     inUser.Email,
		Password:  inUser.Password}

	ur.data[id] = user

	outUser := models.OutUser{
		Id:        user.Id,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	return outUser, nil
}

func (ur *UserRepository) GetUsers() []models.OutUser {
	var outUsers []models.OutUser

	for _, user := range ur.data {
		outUsers = append(outUsers, models.OutUser{
			Id:        user.Id,
			FirstName: user.FirstName,
			Email:     user.Email,
		})
	}

	return outUsers
}

// TODO
func (ur *UserRepository) GetUserById() {
}

// TODO
func (ur *UserRepository) UpdateUserById() {
}

// TODO
func (ur *UserRepository) DeleteUser() {
}
