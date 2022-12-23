package message

import (
	message "msg-app/internal/core/messages/impl"
	serv "msg-app/internal/core/messages/service"
)

func EmailSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.EmailService{})
}

func FileSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&message.FileService{})
}
