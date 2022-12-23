package api

import (
	"msg-app/src/api/controllers/message"

	uc "msg-app/src/api/controllers/user"

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

	api.router.Get("/email", message.HandleEmails)

	api.router.Route("/users", func(r chi.Router) {
		// TODO
		// r.Use(AuthMiddleware)
		uc.SetupUserController(r, api.db)
	})

}

func (api *API) GetRouter() *chi.Mux {
	return api.router
}
