package main

import "fmt"

func main() {
	i := question()

	if i != "y" && i != "n" {
		return
	}

	for i != "n" {
		option := printMenu()
		output(option)
		i = question()
		if i != "y" && i != "n" {
			return
		}
	}

	fmt.Println("Goodbye!")
}

func printMenu() int {
	var option int
	fmt.Println("======Calculator======\n\n1 - '+'\n2 - '-'\n3 - '*'\n4 - '/'")
	fmt.Scanln(&option)

	return option
}

func inputNum() (int, int) {
	var first int
	var second int

	fmt.Println("Enter two numbers:")
	fmt.Scanln(&first, &second)

	return first, second
}

func plus(first, second int) int {
	return first + second
}

func minus(first, second int) int {
	return first - second
}

func myltiple(first int, second int) int {
	return first * second
}

func dilit(first int, second int) float64 {
	return float64(first) / float64(second)
}

func output(option int) {
	first, second := inputNum()
	switch option {

	case 1:
		fmt.Println(first, "+", second, "=", plus(first, second))

	case 2:
		fmt.Println(first, "-", second, "=", minus(first, second))

	case 3:
		fmt.Println(first, "*", second, "=", myltiple(first, second))

	case 4:
		fmt.Println(first, "/", second, "=", dilit(first, second))

	default:
		fmt.Println("Incorrect enter")
	}
}

func question() string {
	var i string
	fmt.Println("Do you want to continue? (y/n)")
	fmt.Scanln(&i)

	return i
}
