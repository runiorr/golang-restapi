package message

import "fmt"

type SMService struct{}

func (s *SMService) Send(to, body string) error {
	fmt.Printf("Send '%s' to %s via SMS\n", body, to)

	return nil
}
