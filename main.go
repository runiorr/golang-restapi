package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"msg-app/api/controllers/message"
	"msg-app/profiler"
)

func main() {
	go profiler.MemoryProfiler()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/email", message.HandleEmails)
	r.Get("/sms", message.HandleSMS)

	http.ListenAndServe(":3000", r)
}
