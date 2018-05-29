package publishers

import (
	"fmt"
	"time"

	"github.com/husainaloos/event-bus/messages"
	"github.com/satori/go.uuid"
)

// DemoPublisher demo publisher for testing and demonistration only
type DemoPublisher struct {
	id               string
	publishToChannel *chan (messages.Message)
	isRunning        bool
}

// NewDemoPublisher constructor
func NewDemoPublisher(ID string) *DemoPublisher {
	return &DemoPublisher{
		id:               ID,
		publishToChannel: nil,
		isRunning:        false,
	}
}

// ID gets the ID
func (p DemoPublisher) ID() string {
	return p.id
}

// PublishTo publishes to channel
func (p *DemoPublisher) PublishTo(channel *chan (messages.Message)) {
	p.publishToChannel = channel
}

// Run starts the publisher
func (p *DemoPublisher) Run() error {
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
			CreatedAt: time.Now().UTC(),
			ID:        id.String(),
			Payload:   fmt.Sprintf("message from %s", p.id),
			Tags:      nil,
			SourceID:  p.id,
		}
	}

	return nil
}

// Stop stops the publisher
func (p *DemoPublisher) Stop() error {
	p.isRunning = false
	return nil
}
