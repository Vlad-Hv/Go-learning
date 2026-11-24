package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Latter struct {
	name     string
	duration time.Duration
}

func main() {
	var wg sync.WaitGroup
	A := Latter{
		name:     "A",
		duration: 200 * time.Millisecond,
	}
	B := Latter{
		name:     "B",
		duration: 150 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	BhandleCtx, Bcancel := context.WithTimeout(ctx, 300*time.Millisecond)
	AhandleCtx, Acancel := context.WithTimeout(ctx, 80*time.Millisecond)
	defer cancel()
	wg.Go(func() {
		handle(BhandleCtx, B, Bcancel)
	})

	wg.Go(func() {
		handle(AhandleCtx, A, Acancel)
	})

	wg.Wait()

	fmt.Println("finished")
}

func handle(ctx context.Context, latter Latter, cancel context.CancelFunc) {

	select {
	case <-ctx.Done():
		fmt.Println(latter.name)
		fmt.Println(ctx.Err())

	case <-time.After(latter.duration):
		fmt.Println(latter.name)
		fmt.Println("ok")
		cancel()
	}

}
