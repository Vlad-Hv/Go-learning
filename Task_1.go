package main

import "fmt"

func main() {
	var age int

	fmt.Println("Write how old are you:")

	_, err := fmt.Scanln(&age)

	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	if age < 16 {
		fmt.Println("\nAccess denied")
	} else if age == 16 || age == 17 {
		fmt.Println("\nLimited access")
	} else {
		fmt.Println("\nFull access")
	}
}
