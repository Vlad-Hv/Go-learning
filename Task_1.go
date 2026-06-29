package main

import "fmt"

func main() {
	first, second := gettingVariables()
	userChooise(first, second)
}

func gettingVariables() (int, int) {
	var first int
	var second int
	fmt.Print("Write two numbers to calculate it: ")
	fmt.Scanln(&first, &second)

	return first, second
}

func plus(first int, second int) int {
	return first + second
}

func minus(first, second int) int {
	return first - second
}

func umnoshit(first, second int) int {
	return first * second
}

func userChooise(first, second int) {
	var option int
	fmt.Println("What you do wanna do with your numbers? '+' - 1, '-' - 2, '*' - 3")
	fmt.Scanln(&option)

	switch option {

	case 1:
		fmt.Println(first, "+", second, "=", plus(first, second))

	case 2:
		fmt.Println(first, "-", second, "=", minus(first, second))

	case 3:
		fmt.Println(first, "*", second, "=", umnoshit(first, second))

	default:
		fmt.Println("Incorrect input")
	}
}
