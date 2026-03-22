package producer

import (
	"fmt"
)

type Producer[T fmt.Stringer] interface {
	Produce()
	Data() <-chan T
	Close()
}
