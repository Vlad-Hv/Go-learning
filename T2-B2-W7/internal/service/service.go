package service

import (
	event "eventHandler/internal/service/events"
	"strings"
	"sync"
	"time"
)

func Conveyor(messagesA []event.Event, messagesB []event.Event, done <-chan struct{}, results chan<- event.Event) {
	var wg sync.WaitGroup

	messages := make(chan event.Event)
	chA := make(chan event.Event)
	chB := make(chan event.Event)

	wg.Go(func() {
		sendA(messagesA, done, chA)
	})
	wg.Go(func() {
		sendB(messagesB, done, chB)
	})

	wg.Go(func() {
		merge(chA, chB, messages, done)
	})

	wg.Go(func() {
		sender(messages, results, done)
	})

	wg.Wait()
}

func sendA(A []event.Event, done <-chan struct{}, messagesA chan<- event.Event) {
	defer close(messagesA)
	for _, message := range A {
		select {
		case <-done:
			return

		case <-time.After(time.Duration(message.Delay) * time.Millisecond):
			select {
			case <-done:
				return

			case messagesA <- message:
			}
		}
	}
}

func sendB(B []event.Event, done <-chan struct{}, messagesB chan<- event.Event) {
	defer close(messagesB)
	for _, message := range B {
		select {
		case <-done:
			return

		case <-time.After(time.Duration(message.Delay) * time.Millisecond):
			select {
			case <-done:
				return

			case messagesB <- message:
			}
		}
	}
}

func merge(chA <-chan event.Event, chB <-chan event.Event, messages chan<- event.Event, done <-chan struct{}) {
	for chA != nil || chB != nil {
		select {
		case message, ok := <-chA:
			if !ok {
				chA = nil
				continue
			}
			select {
			case <-done:
				return

			case messages <- message:
			}

		case message, ok := <-chB:
			if !ok {
				chB = nil
				continue
			}

			select {
			case <-done:
				return

			case messages <- message:
			}
		}
	}

	close(messages)
}

func sender(messages <-chan event.Event, results chan<- event.Event, done <-chan struct{}) {
	defer close(results)
sender:
	for {
		select {
		case <-done:
			return

		case message, ok := <-messages:
			if !ok {
				break sender
			}
			text := strings.ToUpper(message.Text)
			message.Text = text
			select {
			case <-done:
				return

			case results <- message:
			}
		}
	}
}
