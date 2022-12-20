package message

import "msg-app/api/services/message"

func EmailSenderFactory() *message.MessageService {
	return message.NewMessageService(&EmailService{})
}

func SMSSenderFactory() *message.MessageService {
	return message.NewMessageService(&SMService{})
}
