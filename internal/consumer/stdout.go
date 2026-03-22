package consumer

import (
	"context"
	"fmt"
)

type Consumer[T fmt.Stringer] interface {
	Consume()
}

type stdOutConsumer[T fmt.Stringer] struct {
	ctx  context.Context
	name string
	in   <-chan T
}

func NewStdOutConsumer[T fmt.Stringer](ctx context.Context, name string, in <-chan T) Consumer[T] {
	return &stdOutConsumer[T]{
		ctx:  ctx,
		name: name,
		in:   in,
	}
}

func (c *stdOutConsumer[T]) Consume() {
	for {
		select {
		case value, ok := <-c.in:
			if !ok {
				fmt.Printf("consume channel is close for %s\n", c.name)
				return
			}
			fmt.Println(value)
		case <-c.ctx.Done():
			fmt.Printf("%s interrupted by: %s\n", c.name, c.ctx.Err())
			return
		}
	}
}
