package main

import (
	"accounts/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("dongwon")
	account.Deposit(1000)
	fmt.Println(account.Balance())

	// go에서 에러 처리
	err := account.Withdraw(5000)
	if err != nil {
		// 프로세스 강제 종료
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account)
}
