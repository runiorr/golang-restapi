package api

import (
	"msg-app/src/api/controllers/message"

	uc "msg-app/src/api/controllers/user"

	"msg-app/src/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	database db.IDatabase
	Router   *chi.Mux
}

func NewAPI(database db.IDatabase) *API {
	api := &API{
		database: database,
		Router:   chi.NewRouter(),
	}
	api.SetupRouter()
	return api
}

func (api *API) SetupRouter() {
	router := api.Router

	router.Use(middleware.Logger)

	// TODO
	// router.Post("/login", login)
	// router.Post("/signup", login)

	router.Get("/email", message.HandleEmails)

	router.Route("/users", func(r chi.Router) {
		// TODO
		// r.Use(AuthMiddleware)
		uc.SetupUserRoutes(r)
	})

}
