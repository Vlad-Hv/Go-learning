package main

import (
	"fmt"
)

func main() {
	first, second, err := getNumbers()

	if err != nil {
		fmt.Println(err)
		return
	}

	result := diline(first, second)

	fmt.Println(result)
}

func getNumbers() (int, int, error) {
	var first int
	var second int

	fmt.Print("Enter two numbers to diline: ")
	fmt.Scanln(&first, &second)

	if second == 0 {
		return 0, 0, fmt.Errorf("Incorrect number: %d", second)
	}

	return first, second, nil
}

func diline(first, second int) float64 {
	return float64(first) / float64(second)
}
