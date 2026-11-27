package service

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Work(Actx context.Context, cancel context.CancelFunc, isAError bool, isBCNeed bool) {
	var wg sync.WaitGroup
	Bctx, Bcancel := context.WithCancel(Actx)
	Cctx, Ccancel := context.WithCancel(Actx)
	defer Bcancel()
	defer Ccancel()

	wg.Go(func() {
		select {
		case <-time.After(300 * time.Millisecond):
			if isAError {
				cancel()
				fmt.Println("A Stopped with the error")
				return
			}
			fmt.Println("A")
			fmt.Println("successfully finished")

		case <-Actx.Done():
			fmt.Println("A stopped with the error")
			return
		}
	})

	wg.Go(func() {
		select {
		case <-time.After(400 * time.Millisecond):
			fmt.Println("B")
			fmt.Println("successfully finished")

		case <-Bctx.Done():
			fmt.Println("B stopped")
			return
		}
	})

	wg.Go(func() {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("C")
			fmt.Println("successfully finished")

		case <-Cctx.Done():
			fmt.Println("C stopped")
			return
		}
	})

	time.Sleep(100 * time.Millisecond)
	if !isBCNeed {
		Bcancel()
		Ccancel()
	}

	wg.Wait()

}
