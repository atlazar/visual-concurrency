package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/atlazar/visual-concurrency/internal/consumer"
	"github.com/atlazar/visual-concurrency/internal/dto"
	"github.com/atlazar/visual-concurrency/internal/producer"
)

const (
	hangSec = 5
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	w1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	w2 := producer.NewCountProducer(ctx, "producer-2", hangSec*time.Second)

	c1 := consumer.NewStdOutConsumer[dto.Tick](ctx, "consumer-1", w1.Data())
	c2 := consumer.NewStdOutConsumer[dto.Tick](ctx, "consumer-2", w2.Data())

	wg.Go(w1.Produce)
	wg.Go(w2.Produce)
	wg.Go(c1.Consume)
	wg.Go(c2.Consume)

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
