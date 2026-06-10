package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance < usd {
		return errors.New("Недостаточно средств!")
	}
	user.Balance -= usd
	return nil
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника: ", p)
		}
	}()
	user := User{
		Name:    "Mansur",
		Balance: 150,
	}
	err := Pay(&user, 50)
	if err == nil {
		pp.Println("Оплата произведена успешна!", user)
	} else {
		pp.Println("Недостаточно средств на счету!")
	}
}
