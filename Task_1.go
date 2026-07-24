package main

import "fmt"

type Account struct {
	Username string
	Level    int
	Balance  float64
}

func main() {
	account := Account{
		Username: "Vlad",
		Balance:  1000.00,
		Level:    1,
	}
	changeAccount(&account)
	fmt.Println(account)

}

func changeAccount(account *Account) {
	var username string
	var Balance float64

	fmt.Println("Enter new username and amount money you want add to the original balance: ")
	fmt.Scanln(&username, &Balance)
	account.Username = username
	account.Balance += Balance
	account.Level++
}
