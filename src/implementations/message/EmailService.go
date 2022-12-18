package message

import "fmt"

type EmailService struct{}

func (e *EmailService) Send(to, body string) error {
	fmt.Printf("Send '%s' to %s via EMAIL\n", body, to)

	return nil
}
