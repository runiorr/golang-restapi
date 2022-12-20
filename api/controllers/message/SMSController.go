package message

import (
	"fmt"
	"msg-app/api/implementations/message"
	"net/http"
)

func HandleSMS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome (using sms) " + r.Host))

	smsSender := message.SMSSenderFactory()
	err_ := smsSender.Send("john@example.com", "Hello, John!")
	if err_ != nil {
		fmt.Println(err_)
	}
}
