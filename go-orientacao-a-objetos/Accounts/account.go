package accounts

import "fmt"

type Account struct {
	Name          string
	AgencyNumber  int16
	AccountNumber int32
	Balance       float32
}

func (account *Account) ToString() string {
	message := "Name: " + account.Name + "\n" +
		"Agency: " + fmt.Sprintf("%d", account.AgencyNumber) + "\n" +
		"Account Number: " + fmt.Sprintf("%d", account.AccountNumber) + "\n" +
		"Balance " + fmt.Sprintf("%f", account.Balance)
	return message
}

func (account *Account) Withdraw(amount float32) {
	if amount <= account.Balance && amount > 0 {
		account.Balance -= amount
		fmt.Println("\nWithdraw made with success\n")
	} else {
		fmt.Println("\nNot enough balance\n")
	}

}

func (account *Account) Deposit(amount float32) (string, bool) {
	if amount > 0 {
		account.Balance += amount
		return "Deposit made with success.", true
	}
	return "Deposit failed due to value less than zero.", false
}

func (account *Account) TransferTo(receiverAccount *Account, amount float32) bool {
	if account.Balance >= amount && amount > 0 {
		account.Withdraw(amount)
		receiverAccount.Deposit(amount)
		return true
	} else {
		fmt.Println("Balance not enough..")
		return false
	}
}
