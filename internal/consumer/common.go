package consumer

import "fmt"

type Consumer[T fmt.Stringer] interface {
	Consume()
}
