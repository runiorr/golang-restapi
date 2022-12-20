package message

import (
	impl "msg-app/api/impl/message"
	serv "msg-app/api/services/message"
)

func EmailSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&impl.EmailService{})
}

func SMSSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&impl.SMService{})
}

func FileSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&impl.FileService{})
}
