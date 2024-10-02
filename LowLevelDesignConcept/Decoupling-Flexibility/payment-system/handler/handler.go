package handler

import "payment-system/payment"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) MakePayment(pm payment.PaymentMethod, amount float64) error {
	return pm.ProcessPayment(amount)
}
