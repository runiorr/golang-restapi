package message

import (
	"errors"
	"testing"
)

// MockMessageService is a mock implementation of the MessageService interface.
type MockMessageService struct {
	to      string
	message string
	err     error
}

// Send records the message and returns the configured error.
func (m *MockMessageService) Send(to, message string) error {
	m.to = to
	m.message = message
	return m.err
}

func TestMessageService(t *testing.T) {
	tests := []struct {
		name    string
		service *MockMessageService
		to      string
		message string
		wantErr bool
	}{
		{
			name:    "success",
			service: &MockMessageService{},
			to:      "john@example.com",
			message: "Hello, John!",
		},
		{
			name:    "failure",
			service: &MockMessageService{err: errors.New("send error")},
			to:      "john@example.com",
			message: "Hello, John!",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := NewMessageService(tt.service)
			if err := sender.Send(tt.to, tt.message); (err != nil) != tt.wantErr {
				t.Errorf("Sender.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.service.to != tt.to {
				t.Errorf("Sender.Send() to = %v, want %v", tt.service.to, tt.to)
			}
			if tt.service.message != tt.message {
				t.Errorf("Sender.Send() message = %v, want %v", tt.service.message, tt.message)
			}
		})
	}
}
