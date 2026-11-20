package service

import (
	"documentHandler/internal/service/documents"
	"strings"
	"sync"
	"time"
)

type Result struct {
	ID         string
	WordAmount int
}

func HandleDocument(sourceA []documents.Document, sourceB []documents.Document, results chan<- Result, done <-chan struct{}) {
	chA := make(chan documents.Document)
	chB := make(chan documents.Document)
	document := make(chan documents.Document)
	var wg sync.WaitGroup

	wg.Go(func() {
		sendA(sourceA, chA, done)
	})

	wg.Go(func() {
		sendB(sourceB, chB, done)
	})

	wg.Go(func() {
		merge(chA, chB, document, done)
	})

	wg.Go(func() {
		analise(document, results, done)
	})

	wg.Wait()
}

func analise(documents <-chan documents.Document, results chan<- Result, done <-chan struct{}) {
	defer close(results)
	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 2; i++ {
		wg.Go(func() {
			for {
				select {
				case <-done:
					return

				case doc, ok := <-documents:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case <-time.After(time.Duration(doc.HandleDelay) * time.Millisecond):
						var result Result
						result.WordAmount = len(strings.Fields(doc.Text))
						result.ID = doc.ID
						if result.WordAmount == 0 && result.ID == "" {
							return
						}
						select {
						case <-done:
							return

						case results <- result:
						}
					}
				}
			}

		})
	}
}
