package filter

import (
	"github.com/husainaloos/event-bus/message"
)

// AlwaysAllowFilter filter that always says yet
type AlwaysAllowFilter struct {
}

// NewAlwaysAllowFilter constructor
func NewAlwaysAllowFilter() *AlwaysAllowFilter {
	return &AlwaysAllowFilter{}
}

// Allow always allow the message
func (f AlwaysAllowFilter) Allow(m message.Message) bool {
	return true
}
