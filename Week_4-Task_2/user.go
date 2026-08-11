package main

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Username string
	UserAge  int
	Id       uuid.UUID
}

func getUser(name string, age int, id uuid.UUID) User {
	return User{
		Username: name,
		UserAge:  age,
		Id:       id,
	}
}

func getUsername() string {
	var username string

	fmt.Println("Hello, enter your username:")
	fmt.Scanln(&username)
	return username
}

func getUserAge() int {
	var age int
	fmt.Println("Well, enter your age:")
	fmt.Scanln(&age)
	return age
}

func (user User) PrintInfo() {
	fmt.Println("Name:", user.Username)
	fmt.Println("Age:", user.UserAge)
	fmt.Println("ID:", user.Id)
}
