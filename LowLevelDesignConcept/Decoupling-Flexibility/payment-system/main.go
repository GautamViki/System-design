package main

import (
	"payment-system/handler"
	"payment-system/services"
)

func main() {
	cc := services.CreditCard{CardNumber: "sksbfk83rhfdjbv"}
	pp := services.PayPal{Email: "adjah@gmail.com"}
	bt := services.BankTransfer{AccountNumber: "2345678"}

	handler := handler.NewHandler()
	handler.MakePayment(&cc, 231242)
	handler.MakePayment(&pp, 835834)
	handler.MakePayment(&bt, 84398346)
}
