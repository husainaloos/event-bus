package publishers

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"

	"github.com/husainaloos/event-bus/messages"
)

// DemoPublisher demo publisher for testing and demonistration only
type DemoPublisher struct {
	ID               string
	PublishToChannel *chan (messages.Message)
	isRunning        bool
}

// NewDemoPublisher constructor
func NewDemoPublisher(ID string) *DemoPublisher {
	return &DemoPublisher{
		ID:               ID,
		PublishToChannel: nil,
		isRunning:        false,
	}
}

// GetID gets the ID
func (p DemoPublisher) GetID() string {
	return p.ID
}

// PublishTo publishes to channel
func (p *DemoPublisher) PublishTo(channel *chan (messages.Message)) {
	p.PublishToChannel = channel
}

// Start starts the publisher
func (p *DemoPublisher) Start() {
	p.isRunning = true

	for {
		time.Sleep(1 * time.Second)

		if !p.isRunning {
			break
		}

		id, _ := uuid.NewV4()

		*p.PublishToChannel <- messages.Message{
			CreatedAt: time.Now(),
			ID:        id.String(),
			Payload:   fmt.Sprintf("message from %s", p.ID),
			Tags:      nil,
			SourceID:  p.ID,
		}
	}
}

// Stop stops the publisher
func (p *DemoPublisher) Stop() {
	p.isRunning = false
}
