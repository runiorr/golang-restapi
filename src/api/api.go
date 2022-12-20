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

	// router.Post("/login", login)
	// router.Post("/signup", login)

	// Apply auth middleware to only `GET /users/{id}`
	// router.Group(func(r chi.Router) {
	// 	r.Use(AuthMiddleware)
	// 	r.Get("/users/{id}")
	// })

	router.Get("/email", message.HandleEmails)

	router.Route("/users", func(r chi.Router) {
		uc.SetupUserRoutes(r, api.database)
	})

}
