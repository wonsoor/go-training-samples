package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() Paypal {
	return Paypal{}
}

func (p Paypal) Pay(usd int) int {
	fmt.Println("Оплата с помощью Paypal кошелька!")
	fmt.Println("Размер оплаты: ", usd, "USD")
	id := rand.Int()
	return id
}

func (p Paypal) Cancel(id int) {
	fmt.Println("Отмена Paypal-операций! [ID]: ", id)
}
