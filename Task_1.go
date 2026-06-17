package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {

		if i == 7 {
			continue
		} else if i == 15 {
			break
		}

		fmt.Println(i)
	}
}
