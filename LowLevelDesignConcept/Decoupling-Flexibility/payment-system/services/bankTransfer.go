package services

import "fmt"

type BankTransfer struct {
	AccountNumber string
}

func (bt *BankTransfer) ProcessPayment(amount float64) error {
	fmt.Printf("Processing bank transfer payment of $%.2f\n", amount)
	return nil
}
