package controllers

import (
	"github.com/husainaloos/event-bus/filters"
	"github.com/husainaloos/event-bus/publishers"
	"github.com/husainaloos/event-bus/subscribers"
)

// Controller will control messages from publisher and deliver them to subscribers
type Controller interface {
	ID() string
	RegisterPublisher(publishers.Publisher)
	RegisterSubscriber(filters.Filter, subscribers.Subscriber)
	Run()
	Stop()
}
