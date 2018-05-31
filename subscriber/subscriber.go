package subscriber

import (
	"github.com/husainaloos/event-bus/message"
)

// Subscriber the subscriber to a message
type Subscriber interface {
	ID() string
	Subscribe(message.Message)
	Run() error
	GetDoneChannel() chan (message.Message)
	Stop() error
}
