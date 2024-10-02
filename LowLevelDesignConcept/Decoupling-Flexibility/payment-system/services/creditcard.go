package services

import "fmt"

type CreditCard struct {
	CardNumber string
}

func (cc *CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}
