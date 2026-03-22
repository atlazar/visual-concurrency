package dto

import (
	"fmt"
	"time"
)

type Tick struct {
	Worker    string
	Count     int
	Timestamp time.Time
}

func (t Tick) String() string {
	return fmt.Sprintf("%s at %s produce count %v", t.Worker, t.Timestamp.Format(time.DateTime), t.Count)
}
