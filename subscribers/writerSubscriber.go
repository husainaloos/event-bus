package subscribers

import (
	"fmt"
	"io"

	"github.com/husainaloos/event-bus/messages"
)

// WriterSubscriber Demonistration subscriber
type WriterSubscriber struct {
	id          string
	doneChannel chan (messages.Message)
	isRunning   bool
	writer      io.Writer
}

// NewWriterSubscriber constructor
func NewWriterSubscriber(ID string, w io.Writer) *WriterSubscriber {
	return &WriterSubscriber{
		id:          ID,
		doneChannel: make(chan messages.Message),
		isRunning:   false,
		writer:      w,
	}
}

// ID gets the ID
func (s WriterSubscriber) ID() string {
	return s.id
}

// Subscribe Get the message and sends it
func (s *WriterSubscriber) Subscribe(m messages.Message) {
	if !s.isRunning {
		return
	}

	fmt.Fprintf(s.writer, "subscriber ID: %s received message: %v\n", s.id, m)
	//s.doneChannel <- m
}

// Run starts the subscriber
func (s *WriterSubscriber) Run() {
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
