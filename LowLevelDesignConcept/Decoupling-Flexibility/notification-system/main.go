package main

import (
	"notification-system/notifier"
	"notification-system/services"
)

// SendAlert function works with any Notifier.
func SendAlert(n notifier.Notifier, recipient string, message string) {
	n.SendNotification(recipient, message)
}

func main() {
	emailNotifier := services.Email{Sender: "noreply@example.com"}
	smsNotifier := services.SMS{PhoneNumber: "123456789"}

	// Flexibly sending different types of notifications using the same function.
	SendAlert(&emailNotifier, "john@example.com", "Welcome to our service!")
	SendAlert(&smsNotifier, "123456789", "Your code is 1234")
}
