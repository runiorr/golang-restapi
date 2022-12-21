package user

import (
	m "msg-app/src/api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&m.User{})
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(inUser m.InUser) (m.OutUser, error) {
	user := m.User{
		FirstName: inUser.FirstName,
		LastName:  inUser.LastName,
		Email:     inUser.Email,
		Password:  inUser.Password,
	}

	ur.db.Save(&user)

	outUser := m.OutUser{
		FirstName: inUser.FirstName,
		Email:     inUser.Email,
	}

	return outUser, nil
}

func (ur *UserRepository) GetUserById(id string) m.OutUser {
	var user m.User
	ur.db.Find(&user, "id = ?", id)

	outUser := m.OutUser{
		Id:        id,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	return outUser
}

// TODO
func (ur *UserRepository) UpdateUserById() {
}

// TODO
func (ur *UserRepository) DeleteUser() {
}
