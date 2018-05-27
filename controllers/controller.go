package controllers

import (
	"github.com/husainaloos/event-bus/filters"
	"github.com/husainaloos/event-bus/publishers"
	"github.com/husainaloos/event-bus/subscribers"
)

type Controller interface {
	GetID() string
	RegisterPublisher(publishers.Publisher)
	RegisterSubscriber(filters.Filter, subscribers.Subscriber)
	Start()
	Stop()
}
