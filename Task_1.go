package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(numbers)

	numbers = add(numbers)
	fmt.Println("New slice:", numbers)

	numbers, err := deleteByIndex(numbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New slice:", numbers)
}

func add(slice []int) []int {
	var number int
	fmt.Print("Enter number you wanna add: ")
	fmt.Scanln(&number)

	slice = append(slice, number)
	return slice
}

func deleteByIndex(slice []int) ([]int, error) {
	var index int
	fmt.Println("Enter position, which you wanna delete")
	fmt.Scanln(&index)
	if index > len(slice)-1 || index < 0 {
		return slice, errors.New("incorrect index")
	}

	slice = append(slice[:index], slice[index+1:]...)
	return slice, nil
}
