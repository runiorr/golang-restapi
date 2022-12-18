package message

import "go-api-di/src/implementations/message"

func EmailSenderFactory() *MessageService {
	return NewMessageService(&message.EmailService{})
}

func SMSSenderFactory() *MessageService {
	return NewMessageService(&message.SMService{})
}
