package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"

	internal_jwt "msg-app/internal/auth/jwt"
	uc "msg-app/internal/core/users/controller"
	middleware "msg-app/internal/core/users/middleware"
	ur "msg-app/internal/core/users/repository"
	us "msg-app/internal/core/users/service"
)

func SetupUsers(router chi.Router, db *gorm.DB) {
	userRepository := ur.NewUserRepository(db)
	userService := us.NewUserService(*userRepository)
	userControler := uc.NewUserController(*userService)

	router.Post("/logout", userControler.Logout)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(internal_jwt.TokenAuth))
		r.Use(middleware.LoggedInRedirector)

		r.Post("/register", userControler.Register) // POST - /register
		r.Post("/login", userControler.Login)       // POST - /login
	})

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(internal_jwt.TokenAuth))
		r.Use(middleware.UnloggedInRedirector)

		r.Get("/users/profile", userControler.Profile)        // GET - /users/profile
		r.Get("/users/email", userControler.GetUserByEmail)   // GET - /users/
		r.Put("/users/{id}", userControler.UpdateUserById)    // PUT - /users/{id}
		r.Delete("/users/{id}", userControler.DeleteUserById) // DELETE - /users/{id}
	})

}
