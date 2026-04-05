package consumer

import (
	"context"
	"fmt"
)

type funcConsumer[T fmt.Stringer] struct {
	ctx  context.Context
	name string
	in   <-chan T
	f    func(string)
}

func NewFuncConsumer[T fmt.Stringer](ctx context.Context, name string, in <-chan T, f func(string)) Consumer[T] {
	return &funcConsumer[T]{
		ctx:  ctx,
		name: name,
		in:   in,
		f:    f,
	}
}

func (c *funcConsumer[T]) Consume() {
	for {
		select {
		case value, ok := <-c.in:
			if !ok {
				fmt.Printf("%s consume channel is closed\n", c.name)
				return
			}
			c.f(value.String())
		case <-c.ctx.Done():
			fmt.Printf("%s interrupted by: %s\n", c.name, c.ctx.Err())
			return
		}
	}
}
