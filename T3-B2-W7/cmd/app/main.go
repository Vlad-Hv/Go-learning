package main

import (
	"documentHandler/internal/service"
	"documentHandler/internal/service/documents"
	"fmt"
	"sync"
	"time"
)

func main() {
	First := documents.CreateA()
	Second := documents.CreateB()

	results := make(chan service.Result)
	done := make(chan struct{})
	stat := make(map[string]int)
	var wg sync.WaitGroup

	wg.Go(func() {
		service.HandleDocument(First, Second, results, done)
	})

	timeout := time.After(350 * time.Millisecond)

handler:
	for {
		select {
		case <-timeout:
			close(done)
			wg.Wait()
			break handler

		case result, ok := <-results:
			if !ok {
				wg.Wait()
				break handler
			}
			if result.ID == "A1" || result.ID == "A2" {
				stat["A"]++
			} else {
				stat["B"]++
			}
			fmt.Println("ID:", result.ID, "\nWords amount:", result.WordAmount, "\n")
		}
	}

	fmt.Println("A =", stat["A"])
	fmt.Println("B =", stat["B"])
	select {
	case <-done:
		fmt.Println("Stop reason: Timeout")

	default:
		fmt.Println("Stop reason: All Done")
	}
}
