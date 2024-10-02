package services

import "fmt"

type PayPal struct {
	Email string
}

func (pp *PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}
