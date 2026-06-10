package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("Оплата произведена через банк!")
	fmt.Println("Размер оплаты: ", usd, "USD")
	id := rand.Int()
	return id
}

func (b Bank) Cancel(id int) {
	fmt.Println("Отмена банковсокй-операций! [ID]: ", id)
}
