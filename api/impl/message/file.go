package message

import (
	"fmt"
	"os"
)

type FileService struct{}

func (f *FileService) Send(to, message string) error {
	file, err := os.OpenFile(to, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(message); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
