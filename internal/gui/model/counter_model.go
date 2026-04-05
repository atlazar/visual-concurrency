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

type CounterModel struct {
	counterOneHandler func(string)
	counterTwoHandler func(string)
}

func NewCounterModel() *CounterModel {
	return &CounterModel{
		counterOneHandler: func(string) {},
		counterTwoHandler: func(string) {},
	}
}

func (m *CounterModel) GetInitialLabel() string {
	return "not started"
}

func (m *CounterModel) SetCounterOneHandler(h func(string)) {
	m.counterOneHandler = h
}

func (m *CounterModel) SetCounterTwoHandler(h func(string)) {
	m.counterTwoHandler = h
}

func (m *CounterModel) Run() {
	//TODO fixme
	ctx := context.Background()
	p1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	c1 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-1", p1.Data(), m.counterOneHandler)

	p2 := producer.NewCountProducer(ctx, "producer-2", hangSec*time.Second)
	c2 := consumer.NewFuncConsumer[dto.Tick](ctx, "consumer-2", p2.Data(), m.counterTwoHandler)

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
