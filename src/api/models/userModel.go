package models

type InUser struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname, omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname, omitempty"`
	Email     string `json:"email,omitempty" bson:"email, omitempty"`
	Password  string `json:"password,omitempty" bson:"password, omitempty"`
}

type User struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

type OutUser struct {
	Id        string `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	Email     string `json:"email" bson:"email"`
}
