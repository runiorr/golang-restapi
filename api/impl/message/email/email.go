package message

import "fmt"

type EmailService struct{}

func (e *EmailService) Send(to, message string) error {
	fmt.Printf("Send '%s' to %s via EMAIL\n", message, to)

	return nil
}
