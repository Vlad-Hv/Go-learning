package main

import "fmt"

func main() {
	marks := [5]int{2, 5, 8, 1, 10}
	var sum int

	for _, value := range marks {
		sum += value
	}

	avarage := float64(sum) / float64(len(marks))

	fmt.Println(sum, avarage)

}
