package main

import "fmt"

func main() {
	name := "Vlad"
	year := 16
	var language string = "Python"
	newLang := "Go"

	fmt.Print("My name is ")
	fmt.Println(name)

	fmt.Println("I'm ", year, "years old")

	fmt.Println("My previous favorite Development language was ", language)
	fmt.Println("")
	fmt.Println("I have been learning", newLang, "for 3 hours")
}
