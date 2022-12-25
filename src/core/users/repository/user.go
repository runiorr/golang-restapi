package repository

import (
	m "msg-app/src/core/users/model"

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
	ctx := ur.db.Create(&user)
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
// func (ur *UserRepository) UpdateUserById(id string) error {
// 	user = User{1, "old name"}
// 	user2 = User{2, "old name"}

// 	//fist way
// 	db.First(&user)
// 	user.Name = "new name"
// 	db.Save(&user)

// 	//second way
// 	user2.Name = "new name"
// 	db.Model(&user2).Updates(&user2)
// 	return errors.New("A")
// }

func (ur *UserRepository) DeleteUserById(id string) error {
	ctx := ur.db.Delete(&m.User{}, "id = ?", id)
	if err := ctx.Error; err != nil {
		return err
	}
	return nil
}
