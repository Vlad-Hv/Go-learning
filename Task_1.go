package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Gamer struct {
	Name       string
	Level      int
	Role       string
	AmountGold int
	DoneEvents []string
}

type Event struct {
	ID           int
	EventName    string
	Award        int
	MinimalLevel int
	IsDone       bool
}

func main() {
	gamers := getGamers()
	events := getEvents()
	guildStatistic := []string{}

	for i := 0; i < len(events); i++ {
		hero, event, err := getUserAnswer(gamers, events)

		if err != nil {
			fmt.Println(err)
			return
		}

		eventHelper := events[event]
		eventHelper.IsDone = true
		events[event] = eventHelper

		gamers[hero].AmountGold += events[event].Award
		gamers[hero].DoneEvents = append(gamers[hero].DoneEvents, events[event].EventName)
		guildStatistic = append(guildStatistic, events[event].EventName)
		fmt.Println(gamers[hero])
	}

	fmt.Println(guildStatistic)
}

func getGamers() []Gamer {
	gamers := []Gamer{
		{
			Name:       "Hero",
			Level:      2,
			Role:       "Damager",
			AmountGold: 6200,
			DoneEvents: []string{},
		},

		{
			Name:       "Mark",
			Level:      4,
			Role:       "Healer",
			AmountGold: 13990,
			DoneEvents: []string{},
		},

		{
			Name:       "Clara",
			AmountGold: 680,
			Level:      1,
			Role:       "spirit",
			DoneEvents: []string{},
		},
	}

	return gamers
}

func getEvents() map[int]Event {
	var FirstID int = rand.Intn(200) + 101
	var SecondID int = rand.Intn(200) + 101
	var TherdID int = rand.Intn(200) + 101

	events := map[int]Event{
		1: {
			ID:           FirstID,
			EventName:    "Kill the Dragon",
			Award:        1990,
			MinimalLevel: 3,
			IsDone:       false,
		},

		2: {
			ID:           SecondID,
			EventName:    "Find and save the King",
			Award:        2100,
			MinimalLevel: 2,
			IsDone:       false,
		},

		3: {
			ID:           TherdID,
			EventName:    "Steal wallet",
			Award:        500,
			MinimalLevel: 1,
			IsDone:       false,
		},
	}

	return events
}

func getUserAnswer(gamers []Gamer, events map[int]Event) (int, int, error) {
	var heroName string
	var eventNumber int
	var counter int

	for _, hero := range gamers {
		fmt.Println("\n", hero.Name)
	}

	fmt.Print("\nChoose hero name: ")
	fmt.Scanln(&heroName)

	for i := 0; i < len(gamers); i++ {
		if gamers[i].Name == heroName {
			counter++
		}
	}

	if counter == 0 {
		return 0, 0, errors.New("invalid hero name")
	}

	for number, eventInfo := range events {
		if !events[number].IsDone {
			fmt.Println("\n", number, ":", eventInfo)
		} else {
			continue
		}
	}

	fmt.Println("\n\nWell, choose event number: ")
	_, err := fmt.Scanln(&eventNumber)

	if err != nil {
		return 0, 0, fmt.Errorf("event number must be int")
	}

	if eventNumber < 1 || eventNumber > len(events) {
		return 0, 0, errors.New("invalid option")
	}

	if events[eventNumber].IsDone {
		return 0, 0, errors.New("this task alredy taken")
	}
	counter = 0
	for i := 0; i < len(gamers); i++ {
		if gamers[i].Name == heroName {
			counter = i
		}
	}

	isEnough := gamers[counter].Level >= events[eventNumber].MinimalLevel

	if !isEnough {
		return 0, 0, errors.New("level is not high enough")
	}

	return counter, eventNumber, nil
}
