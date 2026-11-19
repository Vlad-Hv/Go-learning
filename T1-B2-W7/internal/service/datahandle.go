package service

import (
	"fmt"
	"orderHandle/internal/storage"
)

func (s *Service) Save(result storage.Result) error {
	if err := s.Storage.Save(result); err != nil {
		return fmt.Errorf("cannot save: %w", err)
	}

	return nil
}

func MakeResult(status string, workerNumber int, id int) storage.Result {
	var Result storage.Result

	Result.ID = id
	Result.WorkerNumber = workerNumber
	Result.Status = status
	/*if isApproved {
		Result.Status = "Approved"
	} else {
		Result.Status = "Cancled"
	}*/

	return Result
}
