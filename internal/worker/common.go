package worker

import (
	"fmt"
)

type Worker[T fmt.Stringer] interface {
	Do()
	Data() <-chan T
}
