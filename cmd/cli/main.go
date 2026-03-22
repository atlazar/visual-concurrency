package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/atlazar/visual-concurrency/internal/dto"
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
	w1 := NewCountWorker(ctx, "worker-1", time.Duration(0))
	wg.Go(w1.Do)
	w2 := NewCountWorker(ctx, "worker-2", hangSec*time.Second)
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

type Worker[T fmt.Stringer] interface {
	Do()
	Data() <-chan T
}

type countWorker struct {
	ctx        context.Context
	name       string
	startDelay time.Duration
	ticks      chan dto.Tick
}

func NewCountWorker(ctx context.Context, name string, startDelay time.Duration) Worker[dto.Tick] {
	return &countWorker{
		ctx:        ctx,
		name:       name,
		startDelay: startDelay,
		ticks:      make(chan dto.Tick),
	}
}

func (w *countWorker) Do() {
	ticker := time.NewTicker(sleepSec * time.Second)
	defer ticker.Stop()

	if w.startDelay.Nanoseconds() > 0 {
		time.Sleep(w.startDelay)
	}

	var t time.Time
	for i := 0; i < countMax; i++ {
		select {
		case t = <-ticker.C:
			select {
			case w.ticks <- dto.Tick{
				Worker:    w.name,
				Timestamp: t,
				Count:     i,
			}:
			default:
				fmt.Printf("%s unable to write count value\n", w.name)
			}
		case <-w.ctx.Done():
			fmt.Printf("%s interrupted by: %v\n", w.name, w.ctx.Err())
			close(w.ticks)
			return
		}
	}
}

func (w *countWorker) Data() <-chan dto.Tick {
	return w.ticks
}
