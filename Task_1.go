package main

import "fmt"

func main() {

	array := getArray()

	fmt.Print(getSum(array))

}

func getSum(array [5]int) int {

	var total int

	for _, number := range array {

		total += number

	}

	return total

}

func getArray() [5]int {

	var array [5]int

	for i := 0; i < 5; i++ {

		fmt.Print("\nEnter ", i+1, " number: ")

		fmt.Scanln(&array[i])

	}

	return array

}
