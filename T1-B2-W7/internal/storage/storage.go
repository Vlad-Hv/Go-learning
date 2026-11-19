package storage

import (
	"fmt"
	"sync"
)

type ResultStorage struct {
	mu      sync.RWMutex
	Results []Result
}

type Statistick struct {
	Results    map[int]Result
	StopReason string
	err        string
}

type Result struct {
	ID           int
	Status       string
	WorkerNumber int
}

func (s *ResultStorage) Save(rs Result) error {
	s.mu.Lock()
	result := copy(rs)

	s.Results = append(s.Results, result)
	s.mu.Unlock()
	return nil
}

func Create() ResultStorage {
	return ResultStorage{}
}

func copy(r Result) Result {
	var result Result

	result.ID = r.ID
	result.Status = r.Status
	result.WorkerNumber = r.WorkerNumber

	return result
}

func (s *ResultStorage) Statistick(reason string /*mu *sync.RWMutex*/, err error) Statistick {
	results := make(map[int]Result)
	defer s.mu.RUnlock()
	s.mu.RLock()
	for i := 0; i < len(s.Results); i++ {
		results[s.Results[i].ID] = s.Results[i]
	}

	errText := fmt.Sprint(err)
	statistick := Statistick{
		Results:    results,
		StopReason: reason,
		err:        errText,
	}

	return statistick

}
