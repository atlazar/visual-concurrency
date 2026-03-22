package consumer

import (
	"context"
	"fmt"
)

type Consumer[T fmt.Stringer] interface {
	DoConsume()
}

type stdOutConsumer[T fmt.Stringer] struct {
	ctx context.Context
	in  <-chan T
}

func NewStdOutConsumer[T fmt.Stringer](ctx context.Context, in <-chan T) Consumer[T] {
	return &stdOutConsumer[T]{
		ctx: ctx,
		in:  in,
	}
}

func (c *stdOutConsumer[T]) DoConsume() {
	for {
		select {
		case value, ok := <-c.in:
			if !ok {
				return
			}
			fmt.Println(value)
		case <-c.ctx.Done():
			return
		}
	}
}
