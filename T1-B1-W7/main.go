package main

import (
	"fmt"
	"sync"
)

type TicketOffice struct {
	mu      sync.RWMutex
	tickets int
	stat    statistic
}

type statistic struct {
	sold     int
	declined int
}

func main() {
	var wg sync.WaitGroup
	ticketOffice := TicketOffice{
		tickets: 30,
	}

	for i := 0; i < 100; i++ {
		wg.Go(func() {
			ticketOffice.Buy()
		})
	}

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			for i := range 20 {
				ticketAmount := ticketOffice.Remaining()
				fmt.Println(i+1, "checker:", ticketAmount)
			}
		})
	}

	wg.Wait()
	fmt.Println("Left:", ticketOffice.tickets)
	fmt.Println("Sold:", ticketOffice.stat.sold)
	fmt.Println("Declined:", ticketOffice.stat.declined)
}

func (o *TicketOffice) Buy() {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.tickets > 0 {
		o.tickets--
		o.stat.sold++
	} else {
		o.stat.declined++
	}
}

func (o *TicketOffice) Remaining() int {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.tickets
}
