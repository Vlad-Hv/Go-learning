package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("\nAuthorise please\n ")

	name, err := nameGetter()

	if err != nil {
		fmt.Println(err)
		return
	}

	age, err := ageGetter()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n\nYour name:", name, "\nYour age:", age)

}

func nameGetter() (string, error) {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if name == "" {
		return "", errors.New("variable name is empty")
	}

	return name, nil
}

func ageGetter() (int, error) {
	var age int
	fmt.Print("\nWell, Enter your age: ")
	fmt.Scanln(&age)

	if age < 0 {
		return 0, fmt.Errorf("Incorrect age: %d", age)
	}

	return age, nil
}
