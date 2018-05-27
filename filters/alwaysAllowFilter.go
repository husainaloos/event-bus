package filters

import (
	"github.com/husainaloos/event-bus/messages"
)

// AlwaysAllowFilter filter that always says yet
type AlwaysAllowFilter struct {
}

// NewAlwaysAllowFilter constructor
func NewAlwaysAllowFilter() *AlwaysAllowFilter {
	return &AlwaysAllowFilter{}
}

// Allow always allow the message
func (f AlwaysAllowFilter) Allow(m messages.Message) bool {
	return true
}
