package api

import (
	messages "msg-app/internal/core/messages/controller"

	users "msg-app/internal/core/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type API struct {
	db     *gorm.DB
	router *chi.Mux
}

func NewAPI(db *gorm.DB) *API {
	api := &API{db: db, router: chi.NewRouter()}
	api.SetupRouter()
	return api
}

func (api *API) SetupRouter() {
	api.router.Use(middleware.Logger)

	// TODO
	// api.router.Post("/login", login)
	// api.router.Post("/signup", login)

	api.router.Get("/email", messages.HandleEmails)

	api.router.Group(func(router chi.Router) {
		router.Route("/users", func(r chi.Router) {
			users.SetupUsers(r, api.db)
		})
	})

}

func (api *API) GetRouter() *chi.Mux {
	return api.router
}
