package user

import (
	"encoding/json"
	"fmt"
	um "msg-app/src/api/models"
	ur "msg-app/src/api/repository/user"
	us "msg-app/src/api/services/user"

	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type UserController struct {
	service us.UserService
}

func newUserController(service us.UserService) *UserController {
	return &UserController{service: service}
}

func SetupUserController(router chi.Router, db *gorm.DB) {
	userRepository := ur.NewUserRepository(db)
	userService := us.NewUserService(*userRepository)
	userControler := newUserController(*userService)

	router.Post("/", userControler.CreateUser)
	router.Get("/", userControler.GetUserByEmail)
	router.Put("/{id}", userControler.UpdateUserById)
	router.Delete("/{id}", userControler.DeleteUserById)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inUser um.InUser
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
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
