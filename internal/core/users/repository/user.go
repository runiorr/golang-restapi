package repository

import (
	m "msg-app/internal/core/users/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Register(registerUser m.RegisterUser) error {
	user := m.User{
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Email:     registerUser.Email,
		Password:  registerUser.Password,
	}
	ctx := ur.db.Save(&user)
	err := ctx.Error
	return err
}

func (ur *UserRepository) GetUserByEmail(email string) (*m.User, error) {
	var user m.User
	ctx := ur.db.First(&user, "email = ?", email)
	err := ctx.Error
	return &user, err
}

// Todo
func (ur *UserRepository) UpdateUserById(id string) string {
	return "2"
}

// Todo
func (ur *UserRepository) DeleteUserById(id string) string {
	return "2"
}
