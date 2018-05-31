package message

import "time"

// Message is what publishers publish to the controller
type Message struct {
	ID        string
	CreatedAt time.Time
	Payload   interface{}
	Tags      map[string]string
}
