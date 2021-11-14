package main

import (
	"fmt"

	"github.com/golang-study/go-orientacao-a-objetos/Accounts/accounts"
)

func main() {
	accountJoselito := accounts.Account{name: "Joselito", agencyNumber: 12345, accountNumber: 33245, balance: 785.12}
	fmt.Println(accountJoselito.ToString())
	fmt.Println()
	accountMaria := accounts.Account{name: "Maria", agencyNumber: 4321, accountNumber: 44532, balance: 200.}
	fmt.Println(accountMaria.ToString())

	fmt.Println("Balance after transfer:")
	accountJoselito.TransferTo(&accountMaria, 500.)

	fmt.Println("Balance Joselito: ", accountJoselito.balance)
	fmt.Println("Balance Maria: ", accountMaria.balance)

}
