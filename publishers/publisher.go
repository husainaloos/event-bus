package publishers

import "github.com/husainaloos/event-bus/messages"

// Publisher publish messages
type Publisher interface {
	ID() string
	PublishTo(*chan (messages.Message))
	Run() error
	Stop() error
}
