package models

type InUser struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
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
