package message

type IMessageService interface {
	Send(string, string) error
}

type MessageService struct {
	service IMessageService
}

func NewMessageService(messageService IMessageService) *MessageService {
	return &MessageService{service: messageService}
}

func (s *MessageService) Send(to, body string) error {
	return s.service.Send(to, body)
}
