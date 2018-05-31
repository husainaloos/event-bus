package publisher

import "github.com/husainaloos/event-bus/message"

// Publisher publish messages
type Publisher interface {
	ID() string
	PublishTo(*chan (message.Message))
	Run() error
	Stop() error
}
