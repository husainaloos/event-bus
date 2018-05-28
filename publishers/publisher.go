package publishers

import "github.com/husainaloos/event-bus/messages"

// Publisher publish messages
type Publisher interface {
	GetID() string
	PublishTo(*chan (messages.Message))
	Start() error
	Stop()
}
