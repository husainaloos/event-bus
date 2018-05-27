package filters

import (
	"github.com/husainaloos/event-bus/messages"
)

// YesFilter filter that always says yet
type YesFilter struct {
}

// NewYesFilter constructor
func NewYesFilter() *YesFilter {
	return &YesFilter{}
}

// Allow always allow the message
func (f YesFilter) Allow(m messages.Message) bool {
	return true
}
