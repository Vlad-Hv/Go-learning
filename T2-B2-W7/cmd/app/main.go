package main

import (
	"eventHandler/internal/service"
	event "eventHandler/internal/service/events"
	"fmt"
	"sync"
	"time"
)

func main() {
	A := event.CreateAFlow()
	B := event.CreateBFlow()
	results := make(chan event.Event)
	done := make(chan struct{})
	statistick := make(map[string]int)
	var wg sync.WaitGroup

	wg.Go(func() {
		service.Conveyor(A, B, done, results)
	})

getter:
	for {
		select {
		case <-time.After(250 * time.Millisecond):
			close(done)
			wg.Wait()
			break getter

		case result, ok := <-results:
			if !ok {
				wg.Wait()
				break getter
			}
			statistick[result.Name]++
			fmt.Println(result.Text)
		}
	}

	select {
	case <-done:
		fmt.Println("Stop reason: timeout")
	default:
		fmt.Println("Stop reason: all done")
	}

	fmt.Println("A source:", statistick["A"])
	fmt.Println("B source:", statistick["B"])
}
