package main

import "fmt"

func main() {
	var DevLanguages [5]string

	for i := 0; i < len(DevLanguages); i++ {
		fmt.Print("\n Write one DevOps Language: ")
		fmt.Scanln(&DevLanguages[i])
	}

	for index, value := range DevLanguages {
		fmt.Println(index+1, value)
	}
}
