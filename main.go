package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"msg-app/api/controllers/message"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/emails", message.HandleEmails)
	r.Get("/smss", message.HandleSMS)

	http.ListenAndServe(":3000", r)
}
