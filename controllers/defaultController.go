package controllers

import (
	"log"

	"github.com/husainaloos/event-bus/filters"
	"github.com/husainaloos/event-bus/messages"
	"github.com/husainaloos/event-bus/publishers"
	"github.com/husainaloos/event-bus/subscribers"
)

// DefaultController default controller that will forward messages from publishers to subscribers
type DefaultController struct {
	id                string
	publishers        []publishers.Publisher
	subscriptionModel map[filters.Filter][]subscribers.Subscriber
	publishChannel    chan (messages.Message)
	stopSignal        chan (bool)
}

// NewDefaultController constructor
func NewDefaultController(ID string) *DefaultController {
	return &DefaultController{
		id:                ID,
		publishChannel:    make(chan messages.Message),
		publishers:        make([]publishers.Publisher, 0),
		subscriptionModel: make(map[filters.Filter][]subscribers.Subscriber),
		stopSignal:        make(chan bool, 1),
	}
}

// ID gets the ID
func (c DefaultController) ID() string {
	return c.id
}

// RegisterPublisher registers a publisher
func (c *DefaultController) RegisterPublisher(p publishers.Publisher) {
	c.publishers = append(c.publishers, p)
}

// RegisterSubscriber regiseter a subscriber with a filter
func (c *DefaultController) RegisterSubscriber(f filters.Filter, s subscribers.Subscriber) {
	var subscriptionList []subscribers.Subscriber

	subscriptionList = c.subscriptionModel[f]
	if subscriptionList == nil {
		subscriptionList = make([]subscribers.Subscriber, 0)
	}

	subscriptionList = append(subscriptionList, s)

	c.subscriptionModel[f] = subscriptionList
}

// Start starts the controller
func (c *DefaultController) Start() {
	for _, p := range c.publishers {
		p.PublishTo(&c.publishChannel)

		go func(p publishers.Publisher) {
			err := p.Run()
			if err != nil {
				log.Fatalf("error occured while starting publisher %s: %v", p.ID(), err)
			}
		}(p)
	}

	for _, v := range c.subscriptionModel {
		for _, s := range v {
			go s.Run()
		}
	}

	go c.handlePublishedMessages()

	<-c.stopSignal
}

func (c *DefaultController) handlePublishedMessages() {

	for {
		select {
		case m := <-c.publishChannel:
			for f, subs := range c.subscriptionModel {
				if f.Allow(m) {
					for _, s := range subs {
						s.Subscribe(m)
					}
				}
			}
		}
	}
}

// Stop stops the controller
func (c *DefaultController) Stop() {
	for _, p := range c.publishers {
		if err := p.Stop(); err != nil {
			log.Fatalf("error occured while stopping publisher %s: %v", p.ID(), err)
		}
	}

	for _, v := range c.subscriptionModel {
		for _, s := range v {
			s.Stop()
		}
	}

	c.stopSignal <- true
}
