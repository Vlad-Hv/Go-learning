package main

import "fmt"

func main() {

	var number int

	fmt.Print("Write number: ")
	fmt.Scanln(&number)

	for i := 1; i <= 10; i++ {
		answer := number * i
		fmt.Println(number, "*", i, "=", answer)
	}
}
