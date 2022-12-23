package users

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	uc "msg-app/internal/core/users/controller"
	ur "msg-app/internal/core/users/repository"
	us "msg-app/internal/core/users/service"
)

func SetupUsers(router chi.Router, db *gorm.DB) {
	userRepository := ur.NewUserRepository(db)
	userService := us.NewUserService(*userRepository)
	userControler := uc.NewUserController(*userService)

	// TODO
	// router.Use(AuthMiddleware)

	router.Post("/", userControler.CreateUser)
	router.Get("/", userControler.GetUserByEmail)

	router.Route("/{id}", func(r chi.Router) {
		// r.Get("/", userControler.GetUserById)
		r.Put("/", userControler.UpdateUserById)
		r.Delete("/", userControler.DeleteUserById)
	})

}
