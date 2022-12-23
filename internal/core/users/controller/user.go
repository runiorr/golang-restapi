package controller

import (
	"encoding/json"
	um "msg-app/internal/core/users/model"
	us "msg-app/internal/core/users/service"

	"net/http"
)

type UserController struct {
	service us.UserService
}

func NewUserController(service us.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inUser um.InUser
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := uc.service.CreateUser(inUser); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("'message':'User created.'"))
}

func (uc *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inUser um.InUser
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	outUser, err := uc.service.GetUserByEmail(inUser.Email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	outUserJson, _ := json.Marshal(outUser)
	w.Write(outUserJson)
}

// Todo
func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
}

// Todo
func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
}
