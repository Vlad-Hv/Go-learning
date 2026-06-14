package main

import "fmt"

func main() {
	var option int

	fmt.Println("\nChoose the option\n\n1 → Play\n2 → Settings\n3 → Exit\n ")
	fmt.Scanln(&option)

	switch option {

	case 1:
		fmt.Println("\nHave a good game!\n ")

	case 2:
		fmt.Println("\nYou are not able to change settings\n ")

	case 3:
		fmt.Println("\nGood bye!\n ")
	}
}
