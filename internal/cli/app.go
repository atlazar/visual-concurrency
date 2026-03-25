package cli

import (
	"context"
	"sync"
	"time"

	"github.com/atlazar/visual-concurrency/internal/consumer"
	"github.com/atlazar/visual-concurrency/internal/dto"
	"github.com/atlazar/visual-concurrency/internal/producer"
)

const (
	hangSec = 5
)

type App struct {
	p1 producer.Producer[dto.Tick]
	p2 producer.Producer[dto.Tick]
	c1 consumer.Consumer[dto.Tick]
	c2 consumer.Consumer[dto.Tick]
}

func NewApp(ctx context.Context) *App {
	p1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	p2 := producer.NewCountProducer(ctx, "producer-2", hangSec*time.Second)

	c1 := consumer.NewStdOutConsumer[dto.Tick](ctx, "consumer-1", p1.Data())
	c2 := consumer.NewStdOutConsumer[dto.Tick](ctx, "consumer-2", p2.Data())
	return &App{
		p1: p1,
		p2: p2,
		c1: c1,
		c2: c2,
	}
}

func (r *App) Run() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Go(func() {
		defer r.p1.Close()
		r.p1.Produce()
	})
	wg.Go(func() {
		defer r.p2.Close()
		r.p2.Produce()
	})
	wg.Go(r.c1.Consume)
	wg.Go(r.c2.Consume)
	return &wg
}
