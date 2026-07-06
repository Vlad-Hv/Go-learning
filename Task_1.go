package main

import (
	"errors"
	"fmt"
)

func main() {
	users := map[string]int{
		"Vlad": 16,
		"Dima": 20,
		"Alex": 25,
	}

	name := getFindingName()
	age, err := findName(users, name)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(age)
}

func getFindingName() string {
	var name string
	fmt.Print("Enter name which you wanna find: ")
	fmt.Scanln(&name)

	return name
}

func findName(users map[string]int, name string) (int, error) {
	age, ok := users[name]

	if !ok {
		return 0, errors.New("there is no users with this name")
	}

	return age, nil
}
