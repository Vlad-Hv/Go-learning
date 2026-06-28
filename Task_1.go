package main

import "fmt"

func main() {
	var studentName [5]string
	var studentAge [5]int
	var studentMark [5]int
	var tries int
	var option int

	for i := 0; i < len(studentName); i++ {
		fmt.Print("\n\nWrite name of ", i+1, "'s student: ")
		fmt.Scanln(&studentName[i])
	}

	fmt.Println("\n\n\nWell next step...")

	for i := 0; i < len(studentAge); i++ {

		fmt.Print("\n\nWrite ", studentName[i], "'s age: ")
		_, err := fmt.Scanln(&studentAge[i])

		if err != nil {
			fmt.Println("Incorrect input...")
			return
		}

		for studentAge[i] <= 15 || studentAge[i] >= 30 {
			tries += 1
			if tries >= 3 {
				fmt.Println("\n\nMore than 3 mistakes - You Blocked")
				return
			}

			fmt.Print("\nIncorrect age... REASON: impossible input... SYSTEM: +1 mistake Try again: ")
			fmt.Scanln(&studentAge[i])
		}
	}

	//Debaged this part of program, working well

	for i := 0; i < 5; i++ {
		fmt.Print("\n\nWrite ", studentName[i], "'s final mark: ")
		_, err := fmt.Scanln(&studentMark[i])

		if err != nil {
			fmt.Println("Incorrect input")
			return
		}

		for studentMark[i] < 1 || studentMark[i] > 10 {
			tries += 1
			if tries > 3 {
				fmt.Println("\n\nMore than 3 mistakes... YOU BLOCKED")
				return
			}

			fmt.Print("Incorrect mark... SYSTEM: +1 mistake... Please, try again: ")
			fmt.Scanln(&studentMark[i])
		}
	}
	fmt.Println("\nDone, You authorised all students...\n ")
	//Debaged new version of proggram, working well

	fmt.Print("\n Do you want continue 1 - yes, Other - no:")
	_, err := fmt.Scanln(&option)

	if err != nil || option != 1 {
		return
	}

	for option != 5 {
		fmt.Println("\n\n\n===== Student Manager =====\n\n1 - Show all students\n2 - Show average grade\n3 - Show best student\n4 - Show adult student\n5 - Exit")
		_, err := fmt.Scanln(&option)

		for err != nil {
			fmt.Println("\n\n\n===== Student Manager =====\n\n1 - Show all students\n2 - Show average grade\n3 - Show best student\n4 - Show adult student\n5 - Exit")
			_, err = fmt.Scanln(&option)
		}

		switch option {

		case 1:
			for i := 0; i < 5; i++ {
				fmt.Println("\n\nName:", studentName[i], "\nAge:", studentAge[i], "\nFinal Mark:", studentMark[i])
			}

		case 2:
			var sum int
			for _, value := range studentMark {
				sum += value
			}
			average := float64(sum) / float64(len(studentMark))
			fmt.Println("\nAverage grade:", average)

		case 3:
			var bestMark int

			for _, value := range studentMark {
				if bestMark < value {
					bestMark = value
				}
			}

			switch {

			case bestMark == studentMark[0]:
				fmt.Println("\n ", studentName[0], bestMark)

			case bestMark == studentMark[1]:
				fmt.Println("\n ", studentName[1], bestMark)

			case bestMark == studentMark[2]:
				fmt.Println("\n ", studentName[2], bestMark)

			case bestMark == studentMark[3]:
				fmt.Println("\n ", studentName[3], bestMark)

			case bestMark == studentMark[4]:
				fmt.Println("\n ", studentName[4], bestMark)

			}

		case 4:

			for i := 0; i < len(studentAge); i++ {
				if studentAge[i] >= 18 {
					fmt.Println("\nName:", studentName[i], " Age:", studentAge[i])
				}
			}

		case 5:

			fmt.Println("\n\n\nGood Bye :)\n ")

		default:

			continue

		}
	}
}

// 1 - creating and initilize 3 arrays (names, ages, marks) with for loops
// 2 - for i != ... exit
// 3 switch with 4 simple bloks

// 1 - creating and initilize 3 arrays (names, ages, marks) with Scanln(&name,&age,&mark)
// 2 - for i != ... exit
// 3 switch with 4 simple bloks
