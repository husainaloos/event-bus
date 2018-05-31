package filter

import "github.com/husainaloos/event-bus/message"

// Filter filter through messages if they are acceptable or not
type Filter interface {
	Allow(message.Message) bool
}
