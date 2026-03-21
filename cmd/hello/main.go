package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	sleepSec = 1
	hangSec  = 5
	countMax = 10
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Go(func() {
		countTillInterrupt(ctx)
	})
	wg.Go(func() {
		time.Sleep(hangSec * time.Second)
		countTillInterrupt(ctx)
	})

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

func countTillInterrupt(ctx context.Context) {
	ticker := time.NewTicker(sleepSec * time.Second)
	defer ticker.Stop()

	var t time.Time
	for i := 0; i < countMax; i++ {
		select {
		case t = <-ticker.C:
			fmt.Printf("At %s counter is %v\n", t.Format(time.DateTime), i)
		case <-ctx.Done():
			fmt.Printf("Interrupted by: %v\n", ctx.Err())
			return
		}
	}
}
