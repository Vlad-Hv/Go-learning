package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Statistic struct {
	CompleteOrders int
	CancellOrders  int
	TotalEarned    float64
}

type Courier struct {
	Name      string
	Balance   float64
	Rating    float64
	IsWorking bool
	IsBlocked bool
	Statistic Statistic
}

func main() {
	earned := 1284.23
	courier := createCourier()
	err := courier.StartWork()
	if err != nil {
		fmt.Println(err)
	}

	err = courier.CompleteOrder(earned)

	if err != nil {
		fmt.Println(err)
	}

	err = courier.CompleteOrder(earned)

	if err != nil {
		fmt.Println(err)
	}

	err = courier.CompleteOrder(0)

	if err != nil {
		fmt.Println(err)
	}

	err = courier.CancellOrder()

	if err != nil {
		fmt.Println(err)
	}

	err = courier.Withdraw(291)
	if err != nil {
		fmt.Println(err)
	}

	err = courier.Withdraw(291321)
	if err != nil {
		fmt.Println(err)
	}

	err = courier.StopWork()

	if err != nil {
		fmt.Println(err)
	}

	courier.Block()
	err = courier.StartWork()

	if err != nil {
		fmt.Println(err)
	}

	courier.PrintStatus()

}
func createCourier() Courier {
	return Courier{
		Name:    "Sofia",
		Balance: 2742.23,
		Rating:  0,
	}
}
func (statistic *Statistic) AddCompletedOrder(earned float64) {
	statistic.CompleteOrders++
	statistic.TotalEarned += earned

}

func (statistic *Statistic) AddCancelledOrder() {
	statistic.CancellOrders++
}

func (statistic *Statistic) TotalOrders() int {
	return statistic.CompleteOrders + statistic.CancellOrders
}

func (courier *Courier) StartWork() error {
	if courier.IsBlocked == true {
		return errors.New("user blocked")
	}

	if courier.IsWorking == true {
		return errors.New("courier is already working")
	}

	courier.IsWorking = true
	return nil
}

func (courier *Courier) StopWork() error {
	if courier.IsWorking == false {
		return errors.New("user is already donnot working")
	}

	courier.IsWorking = false
	return nil
}

func (courier *Courier) CompleteOrder(earned float64) error {
	if courier.IsWorking == false {
		return errors.New("courier is not working")
	}

	if courier.IsBlocked == true {
		return errors.New("courier is blocked")
	}

	if earned <= 0 {
		return errors.New("earn must be more than zero")
	}

	var ratingImprove int = rand.Intn(5) + 1

	courier.Balance += earned
	courier.Statistic.AddCompletedOrder(earned) // I removed err validating in thid method because we are making the same validating in this method(courier)
	courier.Rating += float64(ratingImprove)
	return nil
}

func (courier *Courier) CancellOrder() error {
	if !courier.IsWorking {
		return errors.New("courier must working")
	}

	if courier.IsBlocked {
		return errors.New("courier must not be blocked")
	}

	courier.Statistic.AddCancelledOrder()
	courier.Rating--
	return nil
}

func (courier *Courier) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be more than zero")
	}

	if courier.Balance < amount {
		return errors.New("too big amount")
	}

	if courier.IsBlocked {
		return errors.New("blocked account cannot withdraw money")
	}

	courier.Balance -= amount
	return nil
}

func (courier *Courier) Block() {
	courier.IsBlocked = true
	if courier.IsWorking == true {
		courier.IsWorking = false
	}
}

func (courier *Courier) IsReliable() bool {
	// I don't know, эти условия должны одновремено выполнятся или нет, так что сделаю обработку обоих вариантов. Кай, не серчай
	if courier.Rating >= 4 && courier.Statistic.CompleteOrders > courier.Statistic.CancellOrders && !courier.IsBlocked {
		return true
	}
	return false
}

func (courier *Courier) PrintStatus() {
	fmt.Println("Name:", courier.Name)
	fmt.Println("Balance:", courier.Balance)
	fmt.Println("Rating:", courier.Rating)
	fmt.Println("Is working:", courier.IsWorking)
	fmt.Println("Is Blocked:", courier.IsBlocked)
	fmt.Println("Completed orders:", courier.Statistic.CompleteOrders)
	fmt.Println("Cancelled orders:", courier.Statistic.CancellOrders)
	fmt.Println("Total:", courier.Statistic.TotalEarned)
	fmt.Println("Is courier reliable:", courier.IsReliable())
}
