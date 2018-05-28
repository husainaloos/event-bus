package subscribers

import (
	"log"

	"github.com/husainaloos/event-bus/messages"
)

// DemoSubscriber Demonistration subscriber
type DemoSubscriber struct {
	ID          string
	doneChannel chan (messages.Message)
	isRunning   bool
}

// NewDemoSubscriber constructor
func NewDemoSubscriber(ID string) *DemoSubscriber {
	return &DemoSubscriber{
		ID:          ID,
		doneChannel: make(chan messages.Message),
		isRunning:   false,
	}
}

// GetID gets the ID
func (s DemoSubscriber) GetID() string {
	return s.ID
}

// Subscribe Get the message and sends it
func (s *DemoSubscriber) Subscribe(m messages.Message) {
	if !s.isRunning {
		return
	}

	log.Printf("subscriber %s received message: %s with ID %s\n", s.ID, m.Payload, m.ID)
	//s.doneChannel <- m
}

// Start starts the subscriber
func (s *DemoSubscriber) Start() {
	s.isRunning = true
}

// GetDoneChannel returns the channel that will be filled with processed messages
func (s DemoSubscriber) GetDoneChannel() chan (messages.Message) {
	return s.doneChannel
}

// Stop stops the subscriber
func (s *DemoSubscriber) Stop() {
	s.isRunning = false
}
