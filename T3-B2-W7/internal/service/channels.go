package service

import (
	"documentHandler/internal/service/documents"
	"time"
)

func merge(chA <-chan documents.Document, chB <-chan documents.Document, documents chan<- documents.Document, done <-chan struct{}) {
	defer close(documents)
	for chA != nil || chB != nil {
		select {
		case doc, ok := <-chA:
			if !ok {
				chA = nil
				continue
			}
			select {
			case <-done:
				return

			case documents <- doc:
			}

		case doc, ok := <-chB:
			if !ok {
				chB = nil
				continue
			}

			select {
			case <-done:
				return

			case documents <- doc:
			}

		case <-done:
			return
		}
	}
} // сделать горутину сливающую обе ветки, завершающуюся при отсутсвтии получателя и получающая значения из двух разных каналов и перенаправляющую в общий канал их, а после гортину обрабатыввающую с пулом из 2 работников.
func sendA(sourceA []documents.Document, chA chan<- documents.Document, done <-chan struct{}) {
	defer close(chA)
	for _, doc := range sourceA {
		select {
		case <-done:
			return

		case <-time.After(time.Duration(doc.SourceDelay) * time.Millisecond):
			select {
			case <-done:
				return
			case chA <- doc:
			}
		}
	}
}

func sendB(sourceB []documents.Document, chB chan<- documents.Document, done <-chan struct{}) {
	defer close(chB)
	for _, doc := range sourceB {
		select {
		case <-done:
			return

		case <-time.After(time.Duration(doc.SourceDelay) * time.Millisecond):
			select {
			case <-done:
				return

			case chB <- doc:
			}
		}
	}
}
