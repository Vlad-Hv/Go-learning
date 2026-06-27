package main

import "fmt"

func main() {
	var massive [5]int
	for i := 0; i < len(massive); i++ {
		fmt.Print("\nWhite numer", i+1, ":")
		fmt.Scanln(&massive[i])
	}

	for _, value := range massive {
		fmt.Println(value)
	}
}
