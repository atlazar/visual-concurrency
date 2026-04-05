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
	countOneHandler func(string)
	countTwoHandler func(string)
	countCancel     func()
	countWg         *sync.WaitGroup
	finishHandler   func()
}

func NewCounterModel() *Counter {
	return &Counter{
		countOneHandler: func(string) {},
		countTwoHandler: func(string) {},
		countCancel:     nil,
		countWg:         nil,
	}
}

func (m *Counter) SetCounterOneHandler(h func(string)) {
	m.countOneHandler = h
}

func (m *Counter) SetCounterTwoHandler(h func(string)) {
	m.countTwoHandler = h
}

func (m *Counter) SetFinishHandler(h func()) {
	m.finishHandler = h
}

func (m *Counter) StartCount() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	m.countCancel = cancelFunc

	p1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	c1 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-1", p1.Data(), m.countOneHandler)

	p2 := producer.NewCountProducer(ctx, "producer-2", hangSec*time.Second)
	c2 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-2", p2.Data(), m.countTwoHandler)

	var wg sync.WaitGroup
	m.countWg = &wg

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

	go func() {
		wg.Wait()
		m.finishHandler()
	}()
}

func (m *Counter) StopCount() {
	if m.countCancel != nil {
		m.countCancel()
	}
}
