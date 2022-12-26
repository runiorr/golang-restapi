package api

import (
	"fmt"
	"net/http"

	messages "msg-app/src/core/messages/controller"
	users "msg-app/src/core/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type API struct {
	db         *gorm.DB
	router     *chi.Mux
	httpConfig map[string]string
}

func NewAPI(db *gorm.DB, httpConfig map[string]string) *API {
	api := &API{
		db:         db,
		router:     chi.NewRouter(),
		httpConfig: httpConfig}

	return api
}

func (api *API) SetupRouter() {
	api.router.Use(middleware.Logger)

	// Email test route
	api.router.Get("/email", messages.HandleEmails)

	users.SetupUsers(api.router, api.db)
}

func (api *API) Start() {
	port := api.httpConfig["port"]
	fmt.Printf("App listening at port %s\n", port)

	if err := http.ListenAndServe(port, api.router); err != nil {
		fmt.Println(err)
	}
}
