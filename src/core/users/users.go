package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"

	auth_jwt "msg-app/src/auth/jwt"
	uc "msg-app/src/core/users/controller"
	middleware "msg-app/src/core/users/middleware"
	ur "msg-app/src/core/users/repository"
	us "msg-app/src/core/users/service"
)

func SetupUsers(router chi.Router, db *gorm.DB) {
	userRepository := ur.NewUserRepository(db)
	userService := us.NewUserService(userRepository)
	userControler := uc.NewUserController(userService)

	// Public routes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth_jwt.TokenAuth))
		r.Use(middleware.LoggedInRedirector)

		r.Post("/register", userControler.Register) // POST - /register
		r.Post("/login", userControler.Login)       // POST - /login
	})

	// Private routes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth_jwt.TokenAuth))
		r.Use(middleware.UnloggedInRedirector)

		r.Post("/logout", userControler.Logout)               // POST - /logout
		r.Get("/users/profile", userControler.Profile)        // GET - /users/profile
		r.Get("/users/email", userControler.GetUserByEmail)   // GET - /users/
		r.Put("/users/{id}", userControler.UpdateUserById)    // PUT - /users/{id}
		r.Delete("/users/{id}", userControler.DeleteUserById) // DELETE - /users/{id}
	})

}
