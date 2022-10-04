package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// constructor 만드는 법
// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// receiver
// struct의 첫글자를 따서 소문자로 지어야 한다
// deposit x amount on your account
// *은 복사하지말고 호출한 account를 사용해라, 변경할때
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// exception 없음, go에서 에러 핸들링 방법
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount

	// null none
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// owner return
func (a Account) Owner() string {
	return a.owner
}

// 메서드 오버라이딩
// python에서 __str__과 같음
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account. \nHas: ", a.balance)
}
