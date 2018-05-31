package controller

import (
	"github.com/husainaloos/event-bus/filter"
	"github.com/husainaloos/event-bus/publisher"
	"github.com/husainaloos/event-bus/subscriber"
)

// Controller will control messages from publisher and deliver them to subscribers
type Controller interface {
	ID() string
	AddPublisher(publisher.Publisher)
	AddSubscriber(filter.Filter, subscriber.Subscriber)
	Run()
	Stop()
}
