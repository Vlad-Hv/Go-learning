package main

import "fmt"

func main() {
	firstArray := [4]int{1, 2, 3, 4}
	secondArray := change(firstArray)

	fmt.Println(firstArray, "\n", secondArray)
}

func change(array [4]int) [4]int {
	array[0] = 999
	return array
}
