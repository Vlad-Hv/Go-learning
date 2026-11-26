package service

import (
	"context"
	"errors"
	"fmt"
	"handler/internal/service/line"
	"strings"
	"sync"
	"time"
)

var (
	timeout = errors.New("timeout")
)

func Work(ctx context.Context, lines []line.Lines, source []string, results chan<- error) {
	var wg sync.WaitGroup
	var sendCounter int
	var mu sync.Mutex
	ch := make(chan line.Lines)
	err := make(chan error, 1)
	wg.Go(func() {
		sender(ctx, lines, ch, err)
	})

	for i := 0; i < 2; i++ {
		wg.Go(func() {
			for {
				select {
				case line, ok := <-ch:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						results <- ctx.Err()
						mu.Lock()
						sendCounter++
						mu.Unlock()
						return

					default:
						daughterCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond) //kai, do I have to use cancel(), if daughter ctx done?
						handler(daughterCtx, ctx, line, source)
						cancel()
					}

				case <-ctx.Done():
					results <- ctx.Err()
					mu.Lock()
					sendCounter++
					mu.Unlock()
					return
				}
			}
		})
	}

	wg.Wait()

	switch sendCounter {
	case 0:
		deadline, _ := ctx.Deadline()
		if time.Until(deadline) <= 0 {
			results <- ctx.Err()
			results <- ctx.Err()
			break
		}
		eror := <-err
		results <- eror
		results <- eror

	case 1:
		eror := <-err
		results <- eror
	}
}

func sender(ctx context.Context, lines []line.Lines, ch chan<- line.Lines, err chan<- error) {
	defer close(ch)
	for _, line := range lines {
		select {
		case <-ctx.Done():
			fmt.Println(line.Line, ": not started")

		case ch <- line:
		}
	}
	select {
	case <-ctx.Done():
		err <- timeout

	default:
		err <- nil
	}
}

func handler(daughterCtx context.Context, ctx context.Context, line line.Lines, source []string) {
	select {
	case <-daughterCtx.Done():
		if deadline, _ := ctx.Deadline(); time.Until(deadline) <= 0 {
			fmt.Println(strings.ToLower(line.Line), ": run timeout")
			return
		}
		fmt.Println(strings.ToLower(line.Line), ": check timeout")
		return

	case <-time.After(line.Duration):

		for _, word := range source {
			select {
			case <-daughterCtx.Done():
				if deadline, _ := ctx.Deadline(); time.Until(deadline) <= 0 {
					fmt.Println(strings.ToLower(line.Line), ": run timeout")
					return
				}
				fmt.Println(word, ": check timeout")
				return

			default:
				if word == strings.ToLower(line.Line) {
					fmt.Println(word, ": found")
					return
				}
			}
		}

		fmt.Println(strings.ToLower(line.Line), ": not found")

	}
}
