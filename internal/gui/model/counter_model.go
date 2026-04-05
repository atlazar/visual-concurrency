package model

import (
	"context"
	"sync"
	"time"

	"github.com/atlazar/visual-concurrency/internal/consumer"
	"github.com/atlazar/visual-concurrency/internal/dto"
	"github.com/atlazar/visual-concurrency/internal/producer"
)

type CounterModel struct {
	counterOne string
	counterTwo string
}

func NewCounterModel() *CounterModel {
	return &CounterModel{
		counterOne: "one",
		counterTwo: "two",
	}
}

func (m *CounterModel) CounterOneRef() *string {
	return &m.counterOne
}

func (m *CounterModel) CounterTwoRef() *string {
	return &m.counterTwo
}

func (m *CounterModel) Run() {
	//TODO fixme
	ctx := context.Background()
	p1 := producer.NewCountProducer(ctx, "producer-1", time.Duration(0))
	c1 := consumer.NewStrConsumer[dto.Tick](ctx, "consumer-1", p1.Data(), &m.counterOne)

	var wg sync.WaitGroup
	wg.Go(func() {
		defer p1.Close()
		p1.Produce()
	})
	wg.Go(c1.Consume)
}
