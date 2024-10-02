package payment

type PaymentMethod interface {
	ProcessPayment(amount float64) error
}
