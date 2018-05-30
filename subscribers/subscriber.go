package subscribers

import (
	"github.com/husainaloos/event-bus/messages"
)

// Subscriber the subscriber to a message
type Subscriber interface {
	ID() string
	Subscribe(messages.Message)
	Run() error
	GetDoneChannel() chan (messages.Message)
	Stop() error
}
