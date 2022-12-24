package api

import (
	messages "msg-app/src/core/messages/controller"

	users "msg-app/src/core/users"

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

	// Email test route
	api.router.Get("/email", messages.HandleEmails)

	users.SetupUsers(api.router, api.db)
}

func (api *API) GetRouter() *chi.Mux {
	return api.router
}
