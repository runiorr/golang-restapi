package message

import (
	message "msg-app/src/core/messages/impl"
	serv "msg-app/src/core/messages/service"
)

func EmailSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.EmailService{})
}

func FileSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.FileService{})
}
