package filters

import "github.com/husainaloos/event-bus/messages"

// Filter filter through messages if they are acceptable or not
type Filter interface {
	Allow(messages.Message) bool
}
