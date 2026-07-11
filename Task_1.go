package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	var losestrick int
	var history []string
	/*playTimes,*/ avarageCost, balance, err := calcAvarGameCost()

	if err != nil {
		fmt.Println(err)
		return
	}
	statistics := Statistics()
	stickers := stickerMap()

	for /*i := 0; i < playTimes; i++ */ {
		option, err := printMenuAndGetOption()

		if err != nil {
			fmt.Println(err)
			return
		}

		if option == 5 {
			break
		}

		switch option {

		case 1:
			for {
				option, err = continuePlay()

				if err != nil {
					fmt.Println(err)
					return
				}

				if option == 2 {
					break
				}

				if losestrick >= 5 {
					losestrick, balance, err, history = autoWin(stickers, balance, avarageCost, losestrick, statistics, history)

					if err != nil {
						fmt.Println(err)
						break
					}

					statistics["Total games"] += 1

					if len(history) > 10 {
						history = history[1:]
					}

					continue
				}

				balance, err, losestrick, history = playSlots(stickers, avarageCost, balance, losestrick, statistics, history)

				if err != nil {
					fmt.Println(err)
					return
				}

				statistics["Total games"] += 1

				if len(history) > 10 {
					history = history[1:]
				}
			}

		case 2:
			printBalance(balance)

		case 3:
			printStatistics(statistics)

		case 4:
			printHistory(history)
		}

	}

	fmt.Println("Good Bye!\nSee you later!\nYou get", balance)

}

func printStartMessage() (float64, int, error) {
	var balance float64
	var playTimes int
	fmt.Println("\n   Welcome to casino Vladika\nHere my dreams is happening :)")
	fmt.Print("\n\nEnter your starting balance: ")
	_, err := fmt.Scanln(&balance)
	fmt.Print("\nEnter number of games: ")
	_, err2 := fmt.Scanln(&playTimes)
	fmt.Println("\n(if you will still have money after the ending of this number, then you will be able to continue)")

	switch {
	case err != nil || err2 != nil:
		return 0, 0, errors.New("incorrect input type")

	case balance <= 0:
		return 0, 0, errors.New("balance cannot be less then 1")

	case playTimes <= 0:
		return 0, 0, errors.New("0 times to play")
	}
	return balance, playTimes, nil
}

func calcAvarGameCost() ( /*int*/ float64, float64, error) {
	balance, playTimes, err := printStartMessage()

	if err != nil {
		return /*0*/ 0, 0, fmt.Errorf("cannot working: %w", err)
	}

	fmt.Println("\nAvarage game will cost:", balance/float64(playTimes))

	return /*playTimes, */ balance / float64(playTimes), balance, nil
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
func slotsEasy(stickers map[int]string, avarage float64, balance float64, losestrick int, stat map[string]int, history []string) (float64, error, int, []string) {
	var first int
	var second int
	var therd int

	//statistics := Statistics()
	stat["Easy games"] += 1

	if balance < avarage {
		return balance, fmt.Errorf("\nnot enough money to continue playing, your balance: %.2f", balance), losestrick, history
	}

	first = rand.Intn(2) + 1
	second = rand.Intn(2) + 1
	therd = rand.Intn(2) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	balanceFirst := balance

	if first == second && second == therd {
		balance += avarage
		losestrick = 0
		stat["Wins"] += 1
		message := fmt.Sprintf("Easy: win +%.2f", balance-balanceFirst)
		history = append(history, message)
		fmt.Println("You won, your balance", balance)
	} else {
		balance -= avarage
		message := fmt.Sprintf("Easy: lose -%.2f", balanceFirst-balance)
		history = append(history, message)
		losestrick += 1
		stat["Losses"] += 1
		fmt.Println("Will luck in the next time, your balance", balance)
	}
	return balance, nil, losestrick, history
}

func slotsMedium(stickers map[int]string, balance float64, avarage float64, losestrick int, stat map[string]int, history []string) (float64, error, int, []string) {
	var first int
	var second int
	var therd int
	//statistics := Statistics()
	if balance < avarage {
		return balance, fmt.Errorf("not enough money to continue play, your balance: %.2f", balance), losestrick, history
	}
	first = rand.Intn(3) + 1
	second = rand.Intn(3) + 1
	therd = rand.Intn(3) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	stat["Hard games"] += 1
	balanceFirst := balance

	if first == second && second == therd {
		balance += avarage * 4
		losestrick = 0
		message := fmt.Sprintf("Hard: win +%.2f", balance-balanceFirst)
		history = append(history, message)
		stat["Wins"] += 1
		fmt.Println("\nYou won! Your balance:", balance)
	} else {
		balance -= avarage
		losestrick += 1
		message := fmt.Sprintf("Hard: lose -%.2f", balanceFirst-balance)
		history = append(history, message)
		stat["Losses"] += 1
		fmt.Println("You lost! Will luck in another time", balance)
	}

	return balance, nil, losestrick, history
}

func slotsHard(stickers map[int]string, balance float64, avarage float64, losestrick int, stat map[string]int, history []string) (float64, error, int, []string) {
	var first int
	var second int
	var therd int
	//statistics := Statistics()
	if balance < avarage {
		return balance, fmt.Errorf("not enough money to continue play, your balance: %.2f", balance), losestrick, history
	}
	first = rand.Intn(5) + 1
	second = rand.Intn(5) + 1
	therd = rand.Intn(5) + 1

	fmt.Println(stickers[first], stickers[second], stickers[therd])

	stat["MaxWin games"] += 1
	balanceFirst := balance

	if first == second && second == therd {
		balance += avarage * 9
		losestrick = 0
		message := fmt.Sprintf("MaxWin: win +%.2f", balance-balanceFirst)
		history = append(history, message)
		stat["Wins"] += 1
		fmt.Println("\nYou won! Your balance:", balance)
	} else {
		balance -= avarage
		message := fmt.Sprintf("MaxWin: lose -%.2f", balanceFirst-balance)
		history = append(history, message)
		losestrick += 1
		stat["Losses"] += 1
		fmt.Println("You lost! Will luck in another time", balance)
	}

	return balance, nil, losestrick, history
}

func playSlots(stickers map[int]string, avarage float64, balance float64, losestrick int, statistics map[string]int, history []string) (float64, error, int, []string) {
	choice, err := chooseDificulity()

	if err != nil {
		return 0, err, losestrick, history
	}

	switch choice {
	case 1:
		balance, err, losestrick, history = slotsEasy(stickers, avarage, balance, losestrick, statistics, history)

		if err != nil {
			return balance, err, losestrick, history
		}

	case 2:
		balance, err, losestrick, history = slotsMedium(stickers, balance, avarage, losestrick, statistics, history)

		if err != nil {
			return balance, err, losestrick, history
		}

	case 3:
		balance, err, losestrick, history = slotsHard(stickers, balance, avarage, losestrick, statistics, history)

		if err != nil {
			return balance, err, losestrick, history
		}

	}
	return balance, nil, losestrick, history
}

/*func losestrick(losestrick int) (int, error) {
	//losestrick += 1

	if losestrick == 5 {
		return 0, errors.New("\nlosestrick: Autowin\n ")
	}

	return losestrick, nil
}*/

func autoWin(stickers map[int]string, balance float64, avarage float64, losestrick int, stat map[string]int, history []string) (int, float64, error, []string) {
	autoWinX := autoWinConf()
	option, err := chooseDificulity()
	//statistics := Statistics()
	//historyHelp := historyHelp()
	balanceFirst := balance
	if err != nil {
		return losestrick, balance, err, history
	}

	var first int
	var second int
	var third int

	first = rand.Intn(5) + 1
	second = first
	third = second

	fmt.Println(stickers[first], stickers[second], stickers[third])
	balance = (balance - avarage) + avarage*float64(autoWinX[option])
	fmt.Println("\nYou won! Your balance:", balance)

	switch option {
	case 1:
		message := fmt.Sprintf("Easy: bonus win +%2.f", balance-balanceFirst)
		history = append(history, message)

	case 2:
		message := fmt.Sprintf("Hard: bonus win +%2.f", balance-balanceFirst)
		history = append(history, message)

	case 3:
		message := fmt.Sprintf("MaxWin: bonus win +%2.f", balance-balanceFirst)
		history = append(history, message)
	}
	//message := fmt.Sprintf(historyHelp[option], "%d: bonus win + %2.f", balance-balanceFirst)

	stat["Bonus wins"] += 1

	return 0, balance, nil, history
}

func autoWinConf() map[int]int /*error*/ {
	autoWinX := map[int]int{
		1: 2,
		2: 5,
		3: 10,
	}

	return autoWinX
}

func printMenuAndGetOption() (int, error) {
	var option int
	fmt.Println("\n\n===== CASINO VLADIKA =====\n1. Play slots\n2. Show balance\n3. Show statistics\n4. Show history\n5. Exit\nChoose option:")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid option type")
	}

	if option < 1 || option > 5 {
		return 0, fmt.Errorf("incorrect chose: %d", option)
	}

	return option, nil
}

func printBalance(balance float64) {
	fmt.Printf("Your balance: %.2f $", balance)
}

func Statistics() map[string]int {
	stats := map[string]int{
		"Wins":         0,
		"Losses":       0,
		"Bonus wins":   0,
		"Easy games":   0,
		"Hard games":   0,
		"MaxWin games": 0,
		"Total games":  0,
	}

	return stats
}

/*func historyHelp() map[int]string {
	historyHelpDificulity := map[int]string{
		1: "Easy",
		2: "Hard",
		3: "MaxWin",
	}
	return historyHelpDificulity
}*/

func printStatistics(stat map[string]int) {
	for name, statNumber := range stat {
		fmt.Print("\n", name, ": ", statNumber)
	}
}

func printHistory(history []string) {

	if len(history) <= 0 {
		fmt.Println("History is empty")
	} else {
		for index, result := range history {
			fmt.Println("\n", index, result)
		}
	}
}
