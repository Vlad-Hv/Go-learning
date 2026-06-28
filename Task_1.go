package main

import "fmt"

func main() {
	var Prices [5]int
	var sum int
	var theBiggest int

	for i := 0; i < len(Prices); i++ {
		fmt.Print("\n\nWrite ", i+1, " number: ")
		_, err := fmt.Scanln(&Prices[i])

		if err != nil {
			fmt.Println("Incorrect input")
			return
		}

		for Prices[i] <= 0 {
			fmt.Print("\n\nIncorrect price, please, try again: ")
			fmt.Scanln(&Prices[i])
		}
		sum += Prices[i]
	}

	var average float64 = float64(sum) / float64(len(Prices))

	for i := 0; i < len(Prices); i++ {
		if theBiggest < Prices[i] {
			theBiggest = Prices[i]
		}
	}
	//Kai, this comment for you, I know, that there are may be better options: average := float64(sum) / float64(len(Prices)) and for _, value := range Prices{ .. theBiggest < value .. theBiggest = value, I chose another options just for remember, that may be other options :)

	fmt.Println("\n\n\nShoping Report\n\nPrices:", Prices, "\n\nTotal:", sum, "\n\nAverage:", average, "\n\nMost expensive:", theBiggest)

}
