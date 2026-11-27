package main

import (
	"analisys/internal/service"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	isAError := true
	IsBCNeed := false

	service.Work(ctx, cancel, isAError, IsBCNeed)
	fmt.Println("finished")
}
