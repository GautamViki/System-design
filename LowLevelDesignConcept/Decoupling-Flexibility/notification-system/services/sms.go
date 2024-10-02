package services

import "fmt"

type SMS struct {
	PhoneNumber string
}

func (s *SMS) SendNotification(recipient string, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", recipient, message)
	return nil
}