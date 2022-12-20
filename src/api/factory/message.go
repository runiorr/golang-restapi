package message

import (
	"msg-app/src/api/impl/message"
	serv "msg-app/src/api/services/message"
)

func EmailSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.EmailService{})
}

func FileSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.FileService{})
}
