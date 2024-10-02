package notifier

type Notifier interface {
	SendNotification(recipient, message string) error
}
