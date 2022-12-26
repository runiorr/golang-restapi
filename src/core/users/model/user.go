package model

import (
	"time"

	"gorm.io/gorm"
)

type RegisterUser struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname, omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname, omitempty"`
	Email     string `json:"email,omitempty" bson:"email, omitempty"`
	Password  string `json:"password,omitempty" bson:"password, omitempty"`
}

type LoginUser struct {
	Email    string `json:"email,omitempty" bson:"email, omitempty"`
	Password string `json:"password,omitempty" bson:"password, omitempty"`
}

type InUser struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname, omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname, omitempty"`
	Email     string `json:"email,omitempty" bson:"email, omitempty"`
}

type User struct {
	ID        int64 `gorm:"primary_key;auto_increment"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OutUser struct {
	Id        string `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	Email     string `json:"email" bson:"email"`
}

// TODO: Add hooks

// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
// 	if u.Role == "admin" {
// 	  return errors.New("admin user not allowed to delete")
// 	}
// 	return
//   }
