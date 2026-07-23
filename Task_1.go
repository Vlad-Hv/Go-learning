package main

import "fmt"

func main() {
	var balance int = 500

	changeBalance(&balance)
	fmt.Println(balance)
}

func changeBalance(balance *int) {
	*balance += 250
}
