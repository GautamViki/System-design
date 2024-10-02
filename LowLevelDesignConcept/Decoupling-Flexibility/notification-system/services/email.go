package services

import "fmt"

type Email struct {
	Sender string
}

func (e *Email) SendNotification(recipient, message string) error {
	fmt.Printf("Sending Email to %s: %s\n", recipient, message)
	return nil
}
