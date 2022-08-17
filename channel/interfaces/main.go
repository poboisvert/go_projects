package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int) 
	Withdraw(amount int) error
}

// func main()  {

// 	bmo := NewBMO()
// 	bmo.Deposit(500)
// 	if err := bmo.Withdraw(100); err != nil {
// 		panic(err)
// 	}
// 	currentBalance := bmo.GetBalance()
// 	fmt.Printf("The balance is %d\n", currentBalance)
	
// }

func main()  {
	myAccounts := []IBankAccount{
		NewBMO(),
		NewBTC(),
	}

	for _, account := range myAccounts {
		account.Deposit(1000)
		if err := account.Withdraw(50); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}

		balance := account.GetBalance()
		fmt.Printf("balance is %d\n",balance)
	}
}

// go run main.go bmo.go btc.go