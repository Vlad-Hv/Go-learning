package main

import "fmt"

type User struct {
	Name      string
	Age       int
	Country   string
	IsStudent bool
}

func main() {

	studentBeta := User{
		Name:      "Vlad",
		Age:       16,
		Country:   "Belarus",
		IsStudent: true,
	}

	fmt.Println(studentBeta.Name, "\n", studentBeta.Age, "\n", studentBeta.Country, "\n", studentBeta.IsStudent)

	studentBeta.Age = 20
	studentBeta.IsStudent = false

	fmt.Println(studentBeta.Name, "\n", studentBeta.Age, "\n", studentBeta.Country, "\n", studentBeta.IsStudent)
}
