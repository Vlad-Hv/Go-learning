package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	playTimes, avarageCost, balance, err := calcAvarGameCost()

	if err != nil {
		fmt.Println(err)
		return
	}

	stickers := stickerMap()

	for i := 0; i < playTimes; i++ {

		option, err := continuePlay()

		if err != nil {
			fmt.Println(err)
			return
		}

		if option == 2 {
			break
		}

		balance, err = playSlots(stickers, avarageCost, balance)

		if err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("Good Bye!")

}

func printStartMessage() (float64, int, error) {
	var balance float64
	var playTimes int
	fmt.Println("\n   Welcome to casino Vladika\nHere my dreams is happening :)")
	fmt.Print("\n\nEnter your starting balance: ")
	_, err := fmt.Scanln(&balance)
	fmt.Print("\nEnter number of games: ")
	_, err2 := fmt.Scanln(&playTimes)

	switch {
	case err != nil || err2 != nil:
		return 0, 0, errors.New("incorrect input type")

	case balance <= 0:
		return 0, 0, errors.New("balance cannot be zero")

	case playTimes <= 0:
		return 0, 0, errors.New("0 times to play")
	}
	return balance, playTimes, nil
}

func calcAvarGameCost() (int, float64, float64, error) {
	balance, playTimes, err := printStartMessage()

	if err != nil {
		return 0, 0, 0, fmt.Errorf("cannot working: %w", err)
	}

	fmt.Println("\nAvarage game will cost:", balance/float64(playTimes))

	return playTimes, balance / float64(playTimes), balance, nil
}

func chooseDificulity() (int, error) {
	var choice int
	fmt.Println("Choose the dificulity:\n1. Easy x2\n2. Hard x5\n3. MaxWin x10")
	_, err := fmt.Scanln(&choice)

	if err != nil {
		return 0, errors.New("incorrect input type")
	}

	if choice < 1 || choice > 4 {
		return 0, errors.New("incorrect choice")
	}

	return choice, nil
}

func stickerMap() map[int]string {

	stickers := map[int]string{
		1: "🍎",
		2: "🍒",
		3: "🍑",
		4: "🍉",
		5: "🥝",
	}

	return stickers
}

func continuePlay() (int, error) {
	var option int
	fmt.Println("\n🤑Do you want to continue(y - 1/n - 2)?🤑")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid type")
	}

	if option != 1 && option != 2 {
		return 0, fmt.Errorf("valid input is not: %d", option)
	}

	return option, nil

}
func slotsEasy(stickers map[int]string, avarage float64, balance float64) (float64, error) {
	var first int
	var second int
	var therd int

	if balance < avarage {
		return balance, fmt.Errorf("\nnot enough money to continue playing, your balance: %f", balance)
	}

	first = rand.Intn(2) + 1
	second = rand.Intn(2) + 1
	therd = rand.Intn(2) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	if first == second && second == therd {
		balance += avarage
		fmt.Println("You won, your balance", balance)
	} else {
		balance -= avarage
		fmt.Println("Will luck in the next time, your balance", balance)
	}
	return balance, nil
}

func slotsMedium(stickers map[int]string, balance float64, avarage float64) (float64, error) {
	var first int
	var second int
	var therd int

	if balance < avarage {
		return balance, fmt.Errorf("not enough money to continue play, your balance: %f", balance)
	}
	first = rand.Intn(3) + 1
	second = rand.Intn(3) + 1
	therd = rand.Intn(3) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	if first == second && second == therd {
		balance += avarage * 4
		fmt.Println("\nYou won! Your balance:", balance)
	} else {
		balance -= avarage
		fmt.Println("You lost! Will luck in another time")
	}

	return balance, nil
}

func slotsHard(stickers map[int]string, balance float64, avarage float64) (float64, error) {
	var first int
	var second int
	var therd int

	if balance < avarage {
		return balance, fmt.Errorf("not enough money to continue play, your balance: %f", balance)
	}
	first = rand.Intn(5) + 1
	second = rand.Intn(5) + 1
	therd = rand.Intn(5) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	if first == second && second == therd {
		balance += avarage * 9
		fmt.Println("\nYou won! Your balance:", balance)
	} else {
		balance -= avarage
		fmt.Println("You lost! Will luck in another time")
	}

	return balance, nil
}

func playSlots(stickers map[int]string, avarage float64, balance float64) (float64, error) {
	choice, err := chooseDificulity()

	if err != nil {
		return 0, err
	}

	switch choice {
	case 1:
		balance, err = slotsEasy(stickers, avarage, balance)

		if err != nil {
			return balance, err
		}

	case 2:
		balance, err = slotsMedium(stickers, balance, avarage)

		if err != nil {
			return balance, err
		}

	case 3:
		balance, err = slotsHard(stickers, balance, avarage)

		if err != nil {
			return balance, err
		}

	}
	return balance, nil
}
