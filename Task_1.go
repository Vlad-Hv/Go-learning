package main

import "fmt"

func main() {
	numbers := getSlice()
	fmt.Println(getTotal(numbers))
}

func getSlice() []int {
	var numbers []int
	for {
		var number int
		fmt.Print("\nEnter number to add it: ")
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("you left the adding process")
			break
		}

		numbers = append(numbers, number)
	}
	return numbers
}

func getTotal(numbers []int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total
}
