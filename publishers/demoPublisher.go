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
	publishToChannel *chan (messages.Message)
	isRunning        bool
}

// NewDemoPublisher constructor
func NewDemoPublisher(ID string) *DemoPublisher {
	return &DemoPublisher{
		ID:               ID,
		publishToChannel: nil,
		isRunning:        false,
	}
}

// GetID gets the ID
func (p DemoPublisher) GetID() string {
	return p.ID
}

// PublishTo publishes to channel
func (p *DemoPublisher) PublishTo(channel *chan (messages.Message)) {
	p.publishToChannel = channel
}

// Start starts the publisher
func (p *DemoPublisher) Start() error {
	p.isRunning = true

	for {
		time.Sleep(1 * time.Second)

		if !p.isRunning {
			break
		}

		id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		*p.publishToChannel <- messages.Message{
			CreatedAt: time.Now(),
			ID:        id.String(),
			Payload:   fmt.Sprintf("message from %s", p.ID),
			Tags:      nil,
			SourceID:  p.ID,
		}
	}

	return nil
}

// Stop stops the publisher
func (p *DemoPublisher) Stop() error {
	p.isRunning = false
	return nil
}
