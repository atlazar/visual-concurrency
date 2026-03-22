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

func (p *countProducer) Produce() {
	ticker := time.NewTicker(sleepSec * time.Second)
	defer ticker.Stop()

	if p.startDelay.Nanoseconds() > 0 {
		time.Sleep(p.startDelay)
	}

	var t time.Time
	for i := 0; i < countMax; i++ {
		select {
		case t = <-ticker.C:
			select {
			case p.ticks <- dto.Tick{
				Worker:    p.name,
				Timestamp: t,
				Count:     i,
			}:
			default:
				fmt.Printf("%s unable to write count value\n", p.name)
			}
		case <-p.ctx.Done():
			fmt.Printf("%s interrupted by: %v\n", p.name, p.ctx.Err())
			return
		}
	}
}

func (p *countProducer) Data() <-chan dto.Tick {
	return p.ticks
}

func (p *countProducer) Close() {
	close(p.ticks)
}
