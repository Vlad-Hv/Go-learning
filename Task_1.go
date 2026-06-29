package main

import "fmt"

func main() {
	var language string = "Go"
	StartMessage()
	userName()
	userAge()
	Lang(language)
}

func StartMessage() {
	fmt.Println("Welcome to Go Functions Practice")
}

func userName() {
	var name string
	fmt.Println("Write your name ")
	fmt.Scanln(&name)
	fmt.Println("\nName:", name)
}

func userAge() {
	var age int
	fmt.Println("How old are you")
	fmt.Scanln(&age)
	fmt.Println("\nAge:", age)
}

func Lang(lang string) {
	fmt.Println("Language:", lang)
}
