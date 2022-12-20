package message

import "fmt"

type SMService struct{}

func (s *SMService) Send(to, message string) error {
	fmt.Printf("Send '%s' to %s via SMS\n", message, to)

	return nil
}
