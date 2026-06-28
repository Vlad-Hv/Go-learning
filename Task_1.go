package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5

	integerResult := sum / count
	floatResult := float64(sum) / float64(count)

	fmt.Println("Int:", integerResult, "\nFloat:", floatResult)
}
