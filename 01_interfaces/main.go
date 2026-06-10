package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCtypto()
	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("Burger", 20)
	idPhone := paymentModule.Pay("Phone", 100)
	paymentModule.Pay("Something", 999)
	paymentModule.Cancel(idPhone)
	allinfo := paymentModule.AllInfo()
	pp.Println("Все наши покупки:", allinfo)

	pp.Println("Все наши покупки:", allinfo)
}
