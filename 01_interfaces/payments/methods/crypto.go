package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCtypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптовалютой!")
	fmt.Println("Размер оплаты: ", usd, "USDT")
	id := rand.Int()
	return id
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Отмена крипто-операций! [ID]: ", id)
}
