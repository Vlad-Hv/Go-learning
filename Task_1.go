package main

import "fmt"

func main() {
	var username string
	var age int
	var language string
	var option int
	goal := "Junior Go Backend Developer by October"
	correctAge := true
	correctExp := true
	var userExp int

	fmt.Println("hello buddy, Let's make your profile! (Write 1 if you agree)")
	_, err := fmt.Scanln(&option)

	if err != nil || option != 1 {
		fmt.Println("You are not agree")
		return
	}

	fmt.Println("\nOkay, write me please Your name and language: ")
	fmt.Scanln(&username)
	fmt.Scanln(&language)

	fmt.Print("Well, now write your age: ")
	_, err3 := fmt.Scanln(&age)

	if err3 != nil {
		fmt.Println("Wrong 'age' input")
		correctAge = false
	}

	fmt.Print("\nIf your goal: Junior Go Backend Developer by October. Write - 1: ")
	_, err4 := fmt.Scanln(&option)

	if err4 != nil {
		fmt.Println("Incorrect input")
		return
	} else if option != 1 {
		fmt.Print("\nOkay, wrute your goal: ")
		fmt.Scanln(&goal)
	}

	fmt.Print("\nAnd last moment, What is your study experience: ")
	_, err1 := fmt.Scanln(&userExp)

	if err1 != nil {
		fmt.Println("userExp won't show in the case because of: 'Incorrect Exp input'")
		correctExp = false
	}

	fmt.Println("Okay, choose the option:\n\n1 - Show profile\n2 - Show learning progress\n3 - Show goal\n4 - Exit")
	fmt.Scanln(&option)

	switch option {

	case 1:
		fmt.Println("\n User profile:\nName:", username)
		if !correctAge {
			fmt.Println("\nAge: /unknown age\nLanguage:", language)
		} else {
			fmt.Println("\nAge:", age, "\nLanguage:", language)
		}

	case 2:
		if !correctExp {
			fmt.Println("Unknown exp")
		} else {
			fmt.Println("\nLearning progress:", userExp)
		}

	case 3:
		fmt.Println("\nGoal:", goal)

	case 4:
		fmt.Println("\nGood Bye")

	default:
		fmt.Print("\nUnknow option")
	}
}
