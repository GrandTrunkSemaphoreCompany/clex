package clacks

import (
	"time"
)

type Message struct {
	body        string
	created     time.Time
	source      int
	destination int
	route       []int
}

func (m Message) String() string {
	return m.body
}
