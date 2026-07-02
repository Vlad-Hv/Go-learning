package main

import (
	"errors"
	"fmt"
)

func main() {
	age, err := getAge()

	if err != nil {
		fmt.Println(err)
		return
	}

	ageCheker(age)

}

func getAge() (int, error) {
	var userAge int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)

	if userAge < 0 {
		return 0, errors.New("incorrect age")
	}

	return userAge, nil
}

func ageCheker(age int) {
	if age >= 0 {
		fmt.Println("Age accepted")
	}
}
