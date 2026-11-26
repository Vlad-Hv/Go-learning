package main

import (
	"context"
	"fmt"
	"handler/internal/service"
	"handler/internal/service/line"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Millisecond)
	lines := line.Create()
	source := line.CreateSource()
	resutls := make(chan error, 2)
	defer cancel()
	var wg sync.WaitGroup
	start := time.Since(time.Now())
	wg.Go(func() {
		service.Work(ctx, lines, source, resutls)
	})

	wg.Wait()

	firstErr := <-resutls
	secondErr := <-resutls

	fmt.Println(start)
	if firstErr == nil && secondErr == nil {
		fmt.Println("All done")
	} else {
		fmt.Printf("timeout")
	}

}

/*План

1 создать слайс типа строк и вреемени обработки
2 создать сервисе с исполнительнойц частью из отправителя и пула из двух рабоотников*/
