package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type ResearcherCondition struct {
	Health int
	Power  int
	Energy int
}

type Researcher struct {
	ID             int
	Name           string
	Condition      ResearcherCondition
	IsInExpedition bool
}

type Expedition struct {
	Name         string
	Researchers  []*Researcher
	Resourse     []string //[]map[string]int
	Dificulity   int
	IsDone       bool
	IsWon        bool
	WhenWillStop int
}

func main() {
	researchers := createBasicReseacrchers()
	researcherByID := makeIdResearcherMap(&researchers)
	storage := createStorage()
	var userBalance int
	var cicle int
	var expeditions []Expedition

	for {
		option, err := getMainOption()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 7 {
			break
		}

		switch option {
		case 1:
			for {
				researcherOption, err := researcherManageMenu(&researchers, &researcherByID)

				if err != nil {
					fmt.Println(err)
					continue
				}
				cicle++
				if researcherOption == 4 {
					fmt.Println("You left researcher menu")
					break
				}
			}

		case 2:
			for {
				StorageOption, err := storageMenuUX(&storage)

				if err != nil {
					fmt.Println(err)
					continue
				}
				cicle++
				if StorageOption == 4 {
					fmt.Println("You left storage menu")
					break
				}

			}

		case 3:
			fmt.Println("---Creating Expedition---")
			err = createExpedition(researcherByID, &expeditions, &storage, cicle)

			if err != nil {
				fmt.Println(err)
			}
			cicle++

			//fmt.Println(expeditions)

		case 4:
			fmt.Println("\n\n---Researchers Info---\n ", researchers)
			fmt.Println("\n---Storage Info---\n ", storage)
			fmt.Println("\n---Expedition Info---\n ", expeditions, "\n ")

		case 5:
			fmt.Println("Started a new", cicle, "day")
			cicle++

		case 6:
			fmt.Println(userBalance)

		default:
			fmt.Println("incorrect chosen")
		}

		for index := range expeditions {
			if expeditions[index].WhenWillStop == cicle {
				earnedMoney := expeditions[index].expeditionFinished()
				userBalance += earnedMoney
				expeditions[index].IsDone = true
				if expeditions[index].IsWon {
					fmt.Println("\nExpedition", expeditions[index].Name, "finished.\nResult: Successful!\n ")
				} else {
					fmt.Println("\nExpedition", expeditions[index].Name, "finished.\nResult: not Successful\n ")
				}
			}
		}
	}

	fmt.Println("Bye Bye\n you got", userBalance, "$")
}

func createExpedition(researcherByID map[int]*Researcher, expeditioins *[]Expedition, storage *map[string]int, cicle int) error {
	var expedition Expedition
	expedition.getExpeditionName()
	err := expedition.askIDExpedition(researcherByID)
	if err != nil {
		return fmt.Errorf("cannot make an expedition: %w", err)
	}
	err = expedition.chooseDificulty()

	if err != nil {
		return fmt.Errorf("cannot make an expedition: %w", err)
	}
	_, amountCicles := createMapsAmountOfCiclesAndX()
	expedition.WhenWillStop = cicle + amountCicles[expedition.Dificulity]

	err = expedition.getMaterialAmount(storage)
	if err != nil {
		return fmt.Errorf("cannot make an expedition: %w", err)
	}

	*expeditioins = append(*expeditioins, expedition)
	for index := range expedition.Researchers {
		ID := expedition.Researchers[index].ID
		researcherByID[ID].IsInExpedition = true
	}
	// handle dificulity must be 1-3, other = error. and later make steps in main loop, if dificullity = 1 it needs only 2 steps, 3 dificulity -> 6 steps, and chance 1 to 3 (for example) if player win, he will get x5 of resourse
	return nil
}

func (expedition *Expedition) expeditionFinished() int {
	mapOfWinX, _ := createMapsAmountOfCiclesAndX()
	result := randomNumber(mapOfWinX[expedition.Dificulity])
	moneyPrice := moneyMap()
	for index := range expedition.Researchers {
		expedition.Researchers[index].IsInExpedition = false
		expedition.Researchers[index].Condition.Energy -= 25
	}
	if result {
		expedition.IsWon = true
		return moneyPrice[expedition.Dificulity]
	}

	return 0
}

func moneyMap() map[int]int {
	return map[int]int{
		1: 100,
		2: 500,
		3: 5000,
	}
}

func randomNumber(number int) bool {
	first := rand.Intn(number) + 1
	second := rand.Intn(number) + 1

	if first == second {
		return true
	}

	return false
}

func createMapsAmountOfCiclesAndX() (map[int]int, map[int]int) {
	winX := map[int]int{
		1: 2,
		2: 4,
		3: 10,
	}

	amountCicles := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	return winX, amountCicles
}

func (expedition *Expedition) chooseDificulty() error {
	var chose int
	fmt.Println("Choose the dificulty(1 - easy, 2 - medium, 3 - hard)")
	_, err := fmt.Scanln(&chose)

	if err != nil {
		return errors.New("incorrect type")
	}
	if chose < 1 || chose > 3 {
		return errors.New("invalid chose")
	}

	expedition.Dificulity = chose
	return nil
}

func (expedition *Expedition) getMaterialAmount(storage *map[string]int) error {
	var resourse string
	var amount int
	var resourses []string
	fmt.Println("Well, right now you are able to choose 3 type of resourse, be careful, because if you write something wrong, you will lost 1 of the chance")

	for i := 0; i < 3; i++ {
		fmt.Println("Enter name of", i+1, "resourse and amount, which you wanna take:")
		_, err := fmt.Scanln(&resourse, &amount)

		if err != nil {
			fmt.Println("wrong input, be careful, you have", 3-(i+1), "tries")
			continue
		}

		_, ok := (*storage)[resourse]

		if !ok {
			fmt.Println("wrong input, be careful, you have", 3-(i+1), "tries")
			continue
		}

		if ((*storage)[resourse] - amount) < 0 {
			fmt.Println("not enough resourse in the storage, check your amount. You have", 3-(i+1), "tries")
			continue
		}

		(*storage)[resourse] -= amount
		if (*storage)[resourse] == 0 {
			delete(*storage, resourse)
		}
		result := fmt.Sprintf("%s -> %c", resourse, amount)
		resourses = append(resourses, result)
	}

	if len(resourses) == 0 {
		return errors.New("taken resourse is empty")
	}

	expedition.Resourse = resourses
	return nil
}

func (expedition *Expedition) getExpeditionName() {
	var name string
	fmt.Println("Enter name to your expedition:")
	fmt.Scanln(&name)
	expedition.Name = name
}

func (expedition *Expedition) askIDExpedition(researcherByID map[int]*Researcher) /*[]Researcher*/ error {
	var ID int
	var researchersToExpedition []*Researcher
	for {
		fmt.Println("Enter ID of researcher that you wanna add to expedition(if you wanna stop, write smth):")
		_, err := fmt.Scanln(&ID)

		if err != nil {
			fmt.Println("you stopped write ID")
			break
		}

		_, ok := researcherByID[ID]
		if !ok {
			fmt.Println("cannot find researcher with this ID")
			continue
		}

		if researcherByID[ID].IsInExpedition {
			return errors.New("researcher is already in expedition")
		}

		if researcherByID[ID].Condition.Energy < 10 {
			return errors.New("researcher need to relax and get more energy")
		}
		researchersToExpedition = append(researchersToExpedition, researcherByID[ID])
	}

	if len(researchersToExpedition) == 0 {
		return errors.New("reseaarcher list to expedition is empty")
	}

	expedition.Researchers = researchersToExpedition
	return nil
	//return researchersToExpedition

}

func getMainOption() (int, error) {
	var option int
	fmt.Print("---MAIN MENU---\n1. Manage researchers\n2. Manage storage\n3. Manage expeditions\n4. Show base report\n5. Skip one day\n6. Check balance\n7. Exit\nChoose the option: ")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("invalid option type")
	}

	return option, nil
}

func createStorage() map[string]int {
	return map[string]int{
		"food":     5,
		"medicine": 3,
		"fuel":     7,
		"metal":    6,
		"wood":     10,
	}
}
func getOption() (int, error) {
	var option int
	fmt.Println("---Storage Menu---\n1. Add resourse\n2. Spend some resourse\n3. Check storage report\n4. Exit")
	_, err := fmt.Scanln(&option)
	if err != nil {
		return 0, errors.New("incorrect option type")
	}
	return option, nil
}
func storageMenuUX(storage *map[string]int) (int, error) {
	option, err := getOption()
	if err != nil {
		return 0, err
	}

	switch option {
	case 1:
		err = addResourceToTheStorage(storage)
		if err != nil {
			return 0, err
		}

	case 2:
		err = wasteResourse(storage)
		if err != nil {
			return 0, err
		}

	case 3:
		checkHullStorage(*storage)

	case 4:
		return 4, nil

	default:
		return 0, errors.New("incorect chosen")
	}

	return 0, nil
}
func addResourceToTheStorage(storage *map[string]int) error {
	if storage == nil {
		return errors.New("storage is nil")
	}

	resourse, amount, err := getResourseAmount()

	if err != nil {
		return fmt.Errorf("cannot add resourse: %w", err)
	}

	_, ok := (*storage)[resourse]

	if !ok {
		(*storage)[resourse] = amount
		return nil
	}

	(*storage)[resourse] += amount
	return nil
}

func getResourseAmount() (string, int, error) {
	var resourse string
	var amount int
	fmt.Println("Enter resourse name:")
	fmt.Scanln(&resourse)
	fmt.Println("Enter amount of this resourse:")
	_, err := fmt.Scanln(&amount)

	if err != nil {
		return "", 0, errors.New("incorrect amount type")
	}

	if amount < 0 {
		return "", 0, errors.New("amount must be more than zero")
	}

	return resourse, amount, nil
}

func wasteResourse(storage *map[string]int) error {
	if storage == nil {
		return errors.New("storage is nil")
	}
	resourse, amount, err := getResourseAmount()
	if err != nil {
		return fmt.Errorf("cannot spend some resourse: %w", err)
	}

	_, ok := (*storage)[resourse]

	if !ok {
		return fmt.Errorf("there is no resourse %q", resourse)
	}

	if (*storage)[resourse]-amount < 0 {
		fmt.Println("you spend only", (*storage)[resourse], "\nYou have no more")
		delete((*storage), resourse)
		return nil
	}

	(*storage)[resourse] -= amount
	return nil
}

func checkHullStorage(storage map[string]int) {
	fmt.Println("----Storage----")
	for resourseName, amount := range storage {
		fmt.Println(resourseName, ":", amount)
	}
	fmt.Println()
}

func createBasicReseacrchers() []Researcher {
	return []Researcher{
		/*Researcher*/ {
			ID:   001,
			Name: "John",
			Condition: ResearcherCondition{
				Health: 100,
				Power:  41,
				Energy: 100,
			},
			IsInExpedition: false,
		},

		{
			ID:   002,
			Name: "Dilan",
			Condition: ResearcherCondition{
				Health: 100,
				Power:  65,
				Energy: 100,
			},
			IsInExpedition: false,
		},

		{
			ID:   003,
			Name: "Volodia",
			Condition: ResearcherCondition{
				Health: 100,
				Power:  54,
				Energy: 100,
			},
			IsInExpedition: false,
		},
	}
}

func addNewResearcher(researchers *[]Researcher) error {
	if researchers == nil {
		return errors.New("researcher is nil")
	}
	var name string
	var power int

	ID, err := getID(researchers)
	if err != nil {
		return fmt.Errorf("cannot make new researcher: %w", err)
	}

	fmt.Println("Enter researchers name:")
	fmt.Scanln(&name)
	fmt.Println("Enter power to your researcher:")
	_, err = fmt.Scanln(&power)

	if err != nil {
		return errors.New("invalid power type")
	}

	if power < 0 || power > 100 {
		return errors.New("power must be from 0 to 100")
	}

	researcher := Researcher{
		ID:   ID,
		Name: name,
		Condition: ResearcherCondition{
			Health: 100,
			Power:  power,
			Energy: 100,
		},
		IsInExpedition: false,
	}

	*researchers = append(*researchers, researcher)
	return nil
}

func getID(researchers *[]Researcher) (int, error) {
	var ID int
	fmt.Println("To create new researcher enter please his ID:")
	_, err := fmt.Scanln(&ID)

	if err != nil {
		return 0, errors.New("invalid ID type")
	}

	for _, researcher := range *researchers {
		if researcher.ID == ID {
			return 0, errors.New("this ID already use")
		}
	}
	if ID > 999 || ID < 100 {
		return 0, errors.New("id must include 3 numbers")
	}

	return ID, nil
}

func checkAllResearchers(researchers []Researcher) {
	for _, researcher := range researchers {
		fmt.Println(researcher)
	}
}

func makeIdResearcherMap(researchers *[]Researcher) map[int]*Researcher {
	researchersID := make(map[int]*Researcher)

	for index := range *researchers {
		id := (*researchers)[index].ID
		researchersID[id] = &(*researchers)[index]
	}
	return researchersID
}
func chooseOneResearcher(researchers map[int]*Researcher) (int, error) {
	var id int
	fmt.Println("Enter ID researcher:")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return 0, errors.New("invalid id type")
	}

	_, ok := researchers[id]
	if !ok {
		return 0, errors.New("incorrect id")
	}

	fmt.Println("researcher:", id, ":", researchers[id])
	return id, nil
}

func (researcher *Researcher) changeResearcherStat() error {
	name, power, err := getStatInfo()
	if err != nil {
		return fmt.Errorf("cannot change researcher info: %w", err)
	}

	if researcher.IsInExpedition {
		return errors.New("researcher is in expedition, you are not able to change his ifo and stats")
	}

	researcher.Name = name
	researcher.Condition.Power = power
	return nil
}

func getStatInfo() (string, int, error) {
	var name string
	var power int
	fmt.Println("Enter the new name and power:")
	_, err := fmt.Scanln(&name, &power)
	if err != nil {
		return "", 0, errors.New("incorrect input")
	}

	if power < 0 || power > 100 {
		return "", 0, errors.New("power must be from 0 to 100")
	}

	return name, power, nil
}

func researcherManageMenu(researchers *[]Researcher, researcherByID *map[int]*Researcher) (int, error) {
	var option int
	var id int
	fmt.Println("---researcher menu---\n1. Add new researcher\n2. Check all researchers\n3. Choose one researcher by ID\n4. Delete one by ID\n5. Relax one researcher by ID\n6. Exit")
	_, err := fmt.Scanln(&option)
	if err != nil {
		return 0, errors.New("invalid option type")
	}

	switch option {
	case 1:
		err = addNewResearcher(researchers)
		if err != nil {
			return 0, err
		}
		*researcherByID = makeIdResearcherMap(researchers)
		fmt.Println("Added new researcher")

	case 2:
		checkAllResearchers(*researchers)

	case 3:
		id, err = chooseOneResearcher(*researcherByID)
		if err != nil {
			return 0, err
		}
		changeStat, err := WantToCHangeStat()

		if err != nil {
			return 0, err
		}

		if changeStat {
			err = (*researcherByID)[id].changeResearcherStat()
			if err != nil {
				return 0, err
			}
			fmt.Println("Stat changed succesfully")
		}

	case 4:
		err = deleteResearcher(researchers, researcherByID)
		if err != nil {
			return 0, err
		}
		fmt.Println("Researcher deleted succesfully")

	case 5:
		id, err = chooseOneResearcher(*researcherByID)
		if err != nil {
			return 0, err
		}
		err = (*researcherByID)[id].researcherRelax()
		fmt.Println("Researcher relaxed. +5 energy, his energy:", (*researcherByID)[id].Condition.Energy)

	case 6:
		return 4, nil

	default:
		fmt.Println("incorrect chosen")
	}
	return 0, nil
}

func (researcher *Researcher) researcherRelax() error {
	if researcher == nil {
		return errors.New("researcher is nil")
	}

	researcher.Condition.Energy += 5
	return nil
}
func deleteResearcher(researchers *[]Researcher, researcherByID *map[int]*Researcher) error {
	var ID int
	fmt.Println("Enter id researcher that you wanna delete:")
	_, err := fmt.Scanln(&ID)
	if err != nil {
		return errors.New("invalid id type")
	}

	_, ok := (*researcherByID)[ID]
	if !ok {
		return errors.New("incorrect id")
	}

	if researchers == nil {
		return errors.New("researchers is nil")
	}
	if researcherByID == nil {
		return errors.New("researcherByID is nil")
	}

	if (*researcherByID)[ID].IsInExpedition {
		return errors.New("cannot delete researcher because he is in the expedition")
	}
	var Wellindex int
	delete(*researcherByID, ID)
	for index := range *researchers {
		if (*researchers)[index].ID == ID {
			Wellindex = index
		}
	}
	*researchers = append((*researchers)[:Wellindex], (*researchers)[Wellindex+1:]...)
	*researcherByID = makeIdResearcherMap(researchers)
	return nil
}

func WantToCHangeStat() (bool, error) {
	var answer int
	fmt.Println("Do you want change some stats(1 - yes, 2 - No):")
	_, err := fmt.Scanln(&answer)

	if err != nil {
		return false, errors.New("invalid type")
	}

	if answer != 1 && answer != 2 {
		return false, errors.New("incorrect chosen")
	}
	var result bool

	if answer == 1 {
		result = true
	}
	return result, nil
}

//короче, с утра из всех этих функций созданных собрать меню по управлению исследователям. Решить ято делать со слайс и мап, и при удалении из маап удаляется ли из слайс. И написать функцию по удалению
/*
1 making comand
2 manage sclad
3 make an expedition
4 expedition result
5 check base report
*/
