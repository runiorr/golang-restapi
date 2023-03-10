package controller

import (
	"encoding/json"
	um "msg-app/src/core/users/model"
	us "msg-app/src/core/users/service"

	"net/http"
)

type UserController struct {
	service us.IUserService
}

func NewUserController(service us.IUserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message":"Welcome user!"}`))
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
