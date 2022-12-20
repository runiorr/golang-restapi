package models

type InUser struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type User struct {
	Id        string `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type OutUser struct {
	Id        string `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	Email     string `json:"email" bson:"email"`
}
