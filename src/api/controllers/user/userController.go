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

func SetupUserRoutes(router chi.Router, db *gorm.DB) {
	userRepository := ur.NewUserRepository(db)
	userService := us.NewUserService(*userRepository)
	userControler := NewUserController(*userService)

	router.Post("/", userControler.CreateUser)
	router.Get("/{id}", userControler.GetUserById)
	router.Put("/{id}", userControler.UpdateUserById)
	router.Delete("/{id}", userControler.DeleteUser)
}

func NewUserController(service us.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var inUser um.InUser
	err := decoder.Decode(&inUser)
	if err != nil {
		fmt.Println(err)
	}
	outUser, _ := uc.service.CreateUser(inUser)
	outUserJson, _ := json.Marshal(outUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(outUserJson)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	outUser := uc.service.GetUserById(id)
	outUserJson, _ := json.Marshal(outUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(outUserJson)
}

func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
