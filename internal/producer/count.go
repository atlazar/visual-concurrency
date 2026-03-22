package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/atlazar/visual-concurrency/internal/dto"
)

const (
	sleepSec = 1
	countMax = 10
)

type countProducer struct {
	ctx        context.Context
	name       string
	startDelay time.Duration
	ticks      chan dto.Tick
}

func NewCountProducer(ctx context.Context, name string, startDelay time.Duration) Producer[dto.Tick] {
	return &countProducer{
		ctx:        ctx,
		name:       name,
		startDelay: startDelay,
		ticks:      make(chan dto.Tick),
	}
}

func (w *countProducer) Produce() {
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

func (w *countProducer) Data() <-chan dto.Tick {
	return w.ticks
}
