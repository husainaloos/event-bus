package subscribers

import (
	"fmt"
	"io"

	"github.com/husainaloos/event-bus/messages"
)

// WriterSubscriber Demonistration subscriber
type WriterSubscriber struct {
	ID          string
	doneChannel chan (messages.Message)
	isRunning   bool
	writer      io.Writer
}

// NewWriterSubscriber constructor
func NewWriterSubscriber(ID string, w io.Writer) *WriterSubscriber {
	return &WriterSubscriber{
		ID:          ID,
		doneChannel: make(chan messages.Message),
		isRunning:   false,
		writer:      w,
	}
}

// GetID gets the ID
func (s WriterSubscriber) GetID() string {
	return s.ID
}

// Subscribe Get the message and sends it
func (s *WriterSubscriber) Subscribe(m messages.Message) {
	if !s.isRunning {
		return
	}

	fmt.Fprintf(s.writer, "subscriber ID: %s received message: %v\n", s.ID, m)
	//s.doneChannel <- m
}

// Start starts the subscriber
func (s *WriterSubscriber) Start() {
	s.isRunning = true
}

// GetDoneChannel returns the channel that will be filled with processed messages
func (s WriterSubscriber) GetDoneChannel() chan (messages.Message) {
	return s.doneChannel
}

// Stop stops the subscriber
func (s *WriterSubscriber) Stop() {
	s.isRunning = false
}
