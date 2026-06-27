package main

import "fmt"

func main() {
	var allNumbers [5]int
	result := 0

	for i := 0; i < len(allNumbers); i++ {
		fmt.Print("\nWrite number: ")
		fmt.Scanln(&allNumbers[i])
	}

	for _, number := range allNumbers {
		fmt.Println(result, "+", number, "=", result+number)
		result += number

	}

	fmt.Print(result)
}
