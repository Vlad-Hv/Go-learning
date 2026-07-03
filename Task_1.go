package main

import (
	"errors"
	"fmt"
)

func main() {
	const name string = "Vlad"
	const age int = 16
	const password string = "vlad16ag"

	err := checkAge(name, age)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = checkPassword(password)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Welcome, Vlad :)")

}

func checkPassword(password string) error {
	var userPassword string

	fmt.Print("Enter your password down\n>>> ")
	fmt.Scanln(&userPassword)

	if len(userPassword) < 8 {
		return fmt.Errorf("too short password: %d", len(userPassword))
	}

	if userPassword != password {
		return fmt.Errorf("incorrect password: %q", userPassword)
	}

	return nil
}

func checkName(name string) error {

	var userName string
	fmt.Print("Enter your Name: ")
	fmt.Scanln(&userName)

	if userName == "" {
		return errors.New("the 'name' mustn't be empty")
	}

	if userName != name {
		return errors.New("wrong name")
	}

	return nil
}

func checkAge(name string, age int) error {

	err := checkName(name)

	if err != nil {
		return fmt.Errorf("cannot validate age, reason: %w", err)
	}

	var userAge int
	fmt.Println("Enter your age")
	fmt.Scanln(&userAge)

	if userAge < 0 {
		return fmt.Errorf("incorrect age: %d", userAge)
	}

	if userAge != age {
		return errors.New("incorrect age input")
	}

	return nil

}
