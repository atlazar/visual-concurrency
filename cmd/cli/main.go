package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/atlazar/visual-concurrency/internal/worker"
)

const (
	hangSec = 5
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	w1 := worker.NewCountWorker(ctx, "worker-1", time.Duration(0))
	wg.Go(w1.Do)
	w2 := worker.NewCountWorker(ctx, "worker-2", hangSec*time.Second)
	wg.Go(w2.Do)

	go func() {
		//If all goroutine complete - cancel context to gracefully exit process
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	// Restore system signal handling in case of goroutine hang
	// Next SIGINT will result force exit
	cancel()
	wg.Wait()
	fmt.Println("Exit")
}
