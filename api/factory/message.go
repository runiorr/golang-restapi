package message

import (
	email "msg-app/api/impl/message/email"
	file "msg-app/api/impl/message/file"
	sms "msg-app/api/impl/message/sms"
	serv "msg-app/api/services/message"
)

func EmailSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&email.EmailService{})
}

func SMSSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&sms.SMService{})
}

func FileSenderFactory() *serv.MessageService {
	return serv.NewMessageService(&file.FileService{})
}
