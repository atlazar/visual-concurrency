package consumer

import (
	"context"
	"fmt"
)

type strConsumer[T fmt.Stringer] struct {
	ctx  context.Context
	name string
	in   <-chan T
	out  *string
}

func NewStrConsumer[T fmt.Stringer](ctx context.Context, name string, in <-chan T, out *string) Consumer[T] {
	return &strConsumer[T]{
		ctx:  ctx,
		name: name,
		in:   in,
		out:  out,
	}
}

func (c *strConsumer[T]) Consume() {
	for {
		select {
		case value, ok := <-c.in:
			if !ok {
				fmt.Printf("%s consume channel is closed\n", c.name)
				return
			}
			*c.out = value.String()
		case <-c.ctx.Done():
			fmt.Printf("%s interrupted by: %s\n", c.name, c.ctx.Err())
			return
		}
	}
}
