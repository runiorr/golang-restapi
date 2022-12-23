package controller

import (
	"fmt"
	fact "msg-app/internal/core/messages/factory"
	"net/http"
)

func HandleEmails(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome - after reload (using email) " + r.Host))

	emailSender := fact.EmailSenderFactory()
	err := emailSender.Send("john@example.com", "Hello, John!")
	if err != nil {
		fmt.Println(err)
	}
}
