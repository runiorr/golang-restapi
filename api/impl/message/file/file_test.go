package message

import (
	"fmt"
	"msg-app/api/services/message"
	"os"
	"testing"
)

// MockFileService is a mock implementation of the MessageService interface.
type MockFileService struct {
	to      string
	message string
	err     error
}

func (m *MockFileService) Send(to, message string) error {
	m.to = to
	m.message = message
	return m.err
}

func TestFileSend(t *testing.T) {
	sender := message.NewMessageService(&FileService{})
	fileToWrite := "file_test.txt"
	messageToWrite := "this is a file sending test"

	os.Remove(fileToWrite)
	if err := sender.Send(fileToWrite, messageToWrite); err != nil {
		fmt.Print(err)
	}

	sentMessage, _ := os.ReadFile(fileToWrite)
	if messageToWrite != string(sentMessage) {
		t.Errorf("got %q, wanted %q", sentMessage, messageToWrite)
	}
}
