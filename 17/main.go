package main

import "fmt"

type Account struct {
	balance int64
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func (c *Account) simulate(amount int64) int64 {
	c.balance += amount
	return c.balance
}

func main() {
	account := Account{balance: 1000.0}
	account.simulate(100)
	fmt.Println(account.balance)
}