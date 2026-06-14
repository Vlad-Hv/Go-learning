package main

import "fmt"

func main() {
	var countOfDay int

	fmt.Print("Choose 1-7 day to know: Is it -day or -end: ")
	_, err := fmt.Scanln(&countOfDay)

	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	switch countOfDay {

	case 1, 2, 3, 4, 5:
		fmt.Println("\nIt's week day")

	case 6, 7:
		fmt.Println("\nIt's a weekend")
	default:
		fmt.Println("\nWrite a correct week day count")
	}

}
