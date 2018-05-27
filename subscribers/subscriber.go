package subscribers

import (
	"github.com/husainaloos/event-bus/messages"
)

// Subscriber the subscriber to a message
type Subscriber interface {
	GetID() string
	Subscribe(messages.Message)
	Start()
	GetDoneChannel() chan (messages.Message)
	Stop()
}
