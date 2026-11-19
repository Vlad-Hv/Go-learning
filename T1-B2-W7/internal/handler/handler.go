package handler

import (
	"orderHandle/internal/service"
	order "orderHandle/internal/service/orders"
	"orderHandle/internal/storage"
	"sync"
	"time"
)

func Handle(orders []order.Order, workerAmount int, workTime int, s *service.Service) storage.Statistick {
	var oneTimeStorage storage.ResultStorage
	if err := validate(workerAmount, orders, workTime); err != nil {
		return oneTimeStorage.Statistick("error", err)
	}

	var wg sync.WaitGroup
	var help sync.WaitGroup
	var mu sync.Mutex
	var oneMu sync.Mutex
	ch := make(chan *order.Order)
	errCh := make(chan error)
	finish := make(chan struct{})
	results := make(chan string)
	timer := make(chan struct{})
	success := make(chan struct{})
	done := make(chan struct{})

	wg.Go(func() { sender(orders, ch, done) })

	//создать таймаут и ращработать алгоритм по работе горутин
	timeout := time.After(time.Millisecond * (time.Duration(workTime)))
	for i := range workerAmount {
		wg.Go(func() {
			for order := range ch {
				select {
				case <-done:
					return
				default:
				}

				if err := validateOrder(*order); err != nil {
					select {
					case <-done:
						return

					case errCh <- err:
						continue
					}
				}
				status := s.Logic.Check(order, done)

				if err := validateStatus(status); err != nil {
					continue
				}
				result := service.MakeResult(status, i+1, order.ID)

				if err := s.Save(result); err != nil {
					select {
					case <-done:
						return
					case errCh <- err:
					}
				}
				oneMu.Lock()
				oneTimeStorage.Results = append(oneTimeStorage.Results, result)
				oneMu.Unlock()
			}
		})
	}

	help.Go(func() {
		select {
		case <-timeout:
			mu.Lock()
			select {
			case <-done:
				mu.Unlock()
				//results <- "timeout" // тут же по факту не нужно кидать в рещалт тацмаут. тут либо он закрывате доне на дефолте, либо уже произошла ошикба, которая закрыла доне. и завершение вызывает причину ошибка. это для теюя сообщение, кай

			default:
				close(done)
				mu.Unlock()
				select {
				case results <- "timeout":

				case <-timer:
					return
				}

			}

		case <-finish:
			close(success)
		}
	})

	help.Go(
		func() {
			wg.Wait()
			close(finish)
			close(errCh)
		})

	for err := range errCh {
		mu.Lock()
		select {
		case <-done:
			mu.Unlock()
			close(timer)
			wg.Wait()
			help.Wait()
			return oneTimeStorage.Statistick("error", err)

		default:
			close(done)
			close(timer)
			mu.Unlock()
			wg.Wait()
			help.Wait()
			return oneTimeStorage.Statistick("error", err)
		}

	}

	select {
	case reason := <-results:
		help.Wait()
		return oneTimeStorage.Statistick(reason, nil)

	case <-success:
		help.Wait()
		return oneTimeStorage.Statistick("All done", nil)
	}
}

func sender(orders []order.Order, ch chan<- *order.Order, done <-chan struct{}) {
	defer close(ch)
	for i := 0; i < len(orders); i++ {
		select {
		case <-done:
			return

		case ch <- &orders[i]:
		}
	}

}

/*package handler

import (
	"orderHandle/internal/service"
	order "orderHandle/internal/service/orders"
	"orderHandle/internal/storage"
	"sync"
	"time"
)

func Handle(orders []order.Order, workerAmount int, workTime int, s *service.Service) storage.Statistick {
	if err := validate(workerAmount, orders); err != nil {
		return s.Storage.Statistick("error", err)
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan *order.Order)
	errCh := make(chan error)
	results := make(chan string)
	done := make(chan struct{})

	wg.Go(func() { sender(orders, ch) })

	//создать таймаут и ращработать алгоритм по работе горутин
	timeout := time.After(time.Millisecond * (time.Duration(workTime)))
	for i := range workerAmount {
		wg.Go(func() {
			for order := range ch {
				select {
				case <-done:
					return

				default:
				}

				if err := validateOrder(*order); err != nil {
					errCh <- err
				}
				isApproved := s.Check(*order)

				result := service.MakeResult(isApproved, i+1, order.ID)

				if err := s.Storage.Save(result, &mu); err != nil {
					errCh <- err
				}
			}
		})
	}

	go func() {
		select {
		case <-timeout:
			close(done)
			results <- "timeout"
			return

		case <-errCh:
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		return s.Storage.Statistick("error", err)
	}

	select {
	case reason := <-results:
		return s.Storage.Statistick(reason, nil)

	case <-errCh:
		return s.Storage.Statistick("All done", nil)
	}
}

func sender(orders []order.Order, ch chan<- *order.Order) {
	for i := 0; i < len(orders); i++ {
		ch <- &orders[i]
	}

	close(ch)
}*/
