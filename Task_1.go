package main

import "fmt"

func main() {
	users := map[string]int{
		"Vlad": 16,
		"Dima": 20,
		"Alex": 25,
	}
	printUserss(users)
}

func printUserss(users map[string]int) {
	for name, age := range users {
		fmt.Println(name, age)
	}
}
