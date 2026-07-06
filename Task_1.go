package main

import (
	"errors"
	"fmt"
)

func main() {
	userList := map[string]int{
		"Vlad": 16,
		"Dima": 20,
		"Alex": 25,
	}
	err := addUser(userList)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userList)

	err = deleteUser(userList)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userList)
}

func addUser(users map[string]int) error {
	var key string
	var value int

	fmt.Println("Enter name and age:")
	fmt.Scanln(&key, &value)

	if key == "" || value <= 0 {
		return errors.New("invalid name or age")
	}

	users[key] = value
	return nil
}

func deleteUser(users map[string]int) error {
	var nameForDelete string
	fmt.Println("Enter name to delete:")
	fmt.Scanln(&nameForDelete)

	_, ok := users[nameForDelete]

	if !ok {
		return fmt.Errorf("invalid name: %q", nameForDelete)
	}
	delete(users, nameForDelete)
	return nil
}
