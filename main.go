package main

import (
	"fmt"
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

	fmt.Println("App listening at :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
