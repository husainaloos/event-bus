package messages

import "time"

// Message Object to hold data from publisher to subscriber
type Message struct {
	ID        string
	CreatedAt time.Time
	Payload   interface{}
	Tags      map[string]string
	SourceID  string
}
