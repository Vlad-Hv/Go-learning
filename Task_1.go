package main

import "fmt"

func main() {
	var marks [5]int
	var sum int
	var avarageCount int
	var biggestMark int

	//fmt.Println(len(marks))
	for i := 0; i < len(marks); i++ {
		fmt.Print("\n Write ", i+1, " mark: ")
		_, err := fmt.Scanln(&marks[i])

		//if err != nil || marks[i] < 1 || marks[i] > 10{
		//fmt.Print("Uncorrrect input. Please, write correctly")
		if err != nil {
			fmt.Println("Incorrect input")
			return
		}
		for marks[i] < 1 || marks[i] > 10 {

			fmt.Println("Uncorrrect input. Please, write correctly")
			fmt.Print("\n Write ", i+1, " mark: ")
			_, err = fmt.Scanln(&marks[i])

		}

	}

	for _, value := range marks {
		sum += value
	}

	avarageCount = sum / len(marks)
	biggestMark = 0

	for _, value := range marks {
		if biggestMark < value {
			biggestMark = value
		}
	}

	fmt.Println("\n\n\nStudent REPORT\n\nGrades:\n", marks, "\n\nAvarage:\n", avarageCount, "\n\nHighest grade:\n", biggestMark)
}
