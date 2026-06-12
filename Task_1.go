package main

import "fmt"

func main() {
	const appName = "Go task manager"
	var name string
	var age int
	var price float64
	var isStudent bool

	fmt.Println("Application:", appName)

	fmt.Printf("Type of name %T\n", name)
	fmt.Printf("Type of age %T\n", age)
	fmt.Printf("Type of price %T\n", price)
	fmt.Printf("Type of isStudent %T\n", isStudent)

	fmt.Println("Zero value for string:", name)
	fmt.Println("Zero value for int:", age)
	fmt.Println("Zero value for float64:", price)
	fmt.Println("Zero value for bool:", isStudent)

	name = "Vlad"
	age = 16
	price = 120.50
	isStudent = true

	fmt.Println("Next Task:")

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("is Student:", isStudent)

	fmt.Printf("Memorizing, type of tse isStudent is %T\n", isStudent)

}
