package main

import (
	"errors"
	"fmt"
)

func main() {
	first, err := getFirstNumber()

	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := getSecondNumber()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = options(first, second)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFirstNumber() (int, error) {
	var firstNumber int
	fmt.Print("\nEnter first number: ")
	_, err := fmt.Scanln(&firstNumber)

	if err != nil {
		return 0, errors.New("first number's type must be integer")
	}

	return firstNumber, nil
}

func getSecondNumber() (int, error) {
	var secondNumber int
	fmt.Print("\nEnter second number: ")
	_, err := fmt.Scanln(&secondNumber)

	if err != nil {
		return 0, errors.New("second number's type must be integer")
	}

	return secondNumber, nil
}

func plus(first, second int) int {
	return first + second
}

func multiple(first, second int) int {
	return first * second
}

func minus(first, second int) int {
	return first - second
}

func divide(first, second int) (float64, error) {
	if second == 0 {
		return 0, fmt.Errorf("second number must't be %d", second)
	}

	return float64(first) / float64(second), nil
}

func options(first, second int) error {
	var option string
	fmt.Print("Enter operation: ")
	fmt.Scanln(&option)

	switch {

	case option == "+":
		plus := plus(first, second)
		fmt.Println("Result:", plus)

	case option == "-":
		minus := minus(first, second)
		fmt.Println("Result:", minus)

	case option == "*":
		myltiple := multiple(first, second)
		fmt.Println("Result:", myltiple)

	case option == "/":
		divide, err := divide(first, second)

		if err != nil {
			return fmt.Errorf("cannot calculate because of: %w", err)
		}

		fmt.Println("Result: ", divide)

	default:
		return fmt.Errorf("incorrrect operetion: %q", option)
	}
	return nil

}
