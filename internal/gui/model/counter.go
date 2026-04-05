package model

import (
	"context"
	"sync"
	"time"

	"github.com/atlazar/visual-concurrency/internal/consumer"
	"github.com/atlazar/visual-concurrency/internal/dto"
	"github.com/atlazar/visual-concurrency/internal/producer"
)

const hangSec = 5

type Counter struct {
	oneHandler func(string)
	twoHandler func(string)
}

func NewCounterModel() *Counter {
	return &Counter{
		oneHandler: func(string) {},
		twoHandler: func(string) {},
	}
}

func (m *Counter) GetInitialLabel() string {
	return "not started"
}

func (m *Counter) SetCounterOneHandler(h func(string)) {
	m.oneHandler = h
}

func (m *Counter) SetCounterTwoHandler(h func(string)) {
	m.twoHandler = h
}

func (m *Counter) Run() {
	//TODO fixme
	ctx := context.Background()
	p1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	c1 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-1", p1.Data(), m.oneHandler)

	p2 := producer.NewCountProducer(ctx, "producer-2", hangSec*time.Second)
	c2 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-2", p2.Data(), m.twoHandler)

	var wg sync.WaitGroup
	wg.Go(func() {
		defer p1.Close()
		p1.Produce()
	})
	wg.Go(func() {
		defer p2.Close()
		p2.Produce()
	})
	wg.Go(c1.Consume)
	wg.Go(c2.Consume)
}
