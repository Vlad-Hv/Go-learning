package main

import "fmt"

func main() {
	languages := []string{"Golang", "Python", "Java", "Rust"}
	printLangs(languages)

}

func printLangs(array []string) {
	for _, languages := range array {
		fmt.Print(languages, " ")
	}
}
