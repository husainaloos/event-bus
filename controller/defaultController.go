package controller

import (
	"log"

	"github.com/husainaloos/event-bus/filter"
	"github.com/husainaloos/event-bus/message"
	"github.com/husainaloos/event-bus/publisher"
	"github.com/husainaloos/event-bus/subscriber"
)

// DefaultController default controller that will forward messages from publishers to subscribers
type DefaultController struct {
	id                string
	publishers        []publisher.Publisher
	subscriptionModel map[filter.Filter][]subscriber.Subscriber
	publishChannel    chan (message.Message)
	stopSignal        chan (bool)
}

// NewDefaultController constructor
func NewDefaultController(ID string) *DefaultController {
	return &DefaultController{
		id:                ID,
		publishChannel:    make(chan message.Message),
		publishers:        make([]publisher.Publisher, 0),
		subscriptionModel: make(map[filter.Filter][]subscriber.Subscriber),
		stopSignal:        make(chan bool, 1),
	}
}

// ID gets the ID of the controller
func (c DefaultController) ID() string {
	return c.id
}

// AddPublisher adds a publisher to the controller
func (c *DefaultController) AddPublisher(p publisher.Publisher) {
	c.publishers = append(c.publishers, p)
}

// AddSubscriber adds a subscriber to the controller.
// A subscriber is required to have a filter by which the messages are filtered.
func (c *DefaultController) AddSubscriber(f filter.Filter, s subscriber.Subscriber) {
	var subscriptionList []subscriber.Subscriber

	subscriptionList = c.subscriptionModel[f]
	if subscriptionList == nil {
		subscriptionList = make([]subscriber.Subscriber, 0)
	}

	subscriptionList = append(subscriptionList, s)
	c.subscriptionModel[f] = subscriptionList
}

// Start starts the controller
func (c *DefaultController) Start() {
	for _, p := range c.publishers {
		p.PublishTo(&c.publishChannel)

		go func(p publisher.Publisher) {
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
			if err := s.Stop(); err != nil {
				log.Fatalf("error occured while stopping subscriber %s: %v", s.ID(), err)
			}
		}
	}

	c.stopSignal <- true
}
