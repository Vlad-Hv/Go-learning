package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}
	printNum(array)

}

func printNum(array [3]int) {

	for _, namber := range array {
		fmt.Print(namber, " ")
	}
}
