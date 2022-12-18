package main

import (
	"fmt"

	"go-api-di/src/services/message"
)

func main() {
	email_sender := message.EmailSenderFactory()

	err := email_sender.Send("john@example.com", "Hello, John!")
	if err != nil {
		fmt.Println(err)
	}

	sms_sender := message.SMSSenderFactory()

	err_ := sms_sender.Send("john@example.com", "Hello, John!")
	if err_ != nil {
		fmt.Println(err_)
	}
}
