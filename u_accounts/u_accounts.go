package u_accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var NoMoney = errors.New("Can't withdraw")

// new Acocunt create
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// deposit * amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("input deposit : ", amount)
	a.balance = a.balance + amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw you are poor")

		return NoMoney
	}
	a.balance = a.balance - amount
	return nil
}

// change onwer of the account
func (a *Account) ChnageOwner(newOwner string) {
	a.owner = newOwner
}

// onwer of the account

func (a *Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nhas : ", a.Balance())
}
