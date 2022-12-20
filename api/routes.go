package api

import (
	"msg-app/api/controllers/message"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux) {
	router.Get("/email", message.HandleEmails)
	router.Get("/sms", message.HandleSMS)
}
