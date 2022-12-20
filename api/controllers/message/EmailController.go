package message

import (
	"fmt"
	"msg-app/api/implementations/message"
	"net/http"
)

func HandleEmails(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome - after reload (using email) " + r.Host))

	emailSender := message.EmailSenderFactory()
	err := emailSender.Send("john@example.com", "Hello, John!")
	if err != nil {
		fmt.Println(err)
	}
}
