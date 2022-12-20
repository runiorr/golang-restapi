package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"msg-app/api"
	"msg-app/profiler"
)

func main() {
	go profiler.MemoryProfiler()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	api.SetupRoutes(router)
	fmt.Println("App listening at :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
