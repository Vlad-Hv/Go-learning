package main

import "fmt"

type Statistic struct {
	Wins        int
	Losses      int
	GamesPlayed int
}

type GameAccount struct {
	Username  string
	Level     int
	Statistic Statistic
}

func main() {
	statistic := Statistic{
		Wins:        10,
		Losses:      5,
		GamesPlayed: 15,
	}
	gameAccount := GameAccount{
		Username:  "Vlad",
		Level:     7,
		Statistic: statistic,
	}

	fmt.Println("First version:", gameAccount)
	gameAccount = changeStatistic(gameAccount)
	fmt.Println("\nSecond version:", gameAccount)
}

func changeStatistic(account GameAccount) GameAccount {
	account.Statistic.Wins += 1
	account.Statistic.GamesPlayed += 1

	return account
}
