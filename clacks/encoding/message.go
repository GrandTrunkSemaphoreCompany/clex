package encoding

import (
	"time"
)

type Message struct {
	Body        string
	Created     time.Time
	Source      int
	Destination int
	Route       []int
}

func (m Message) String() string {
	return m.Body
}
