package publisher

import (
	"fmt"
	"time"

	"github.com/husainaloos/event-bus/message"
	"github.com/satori/go.uuid"
)

// TimedPublisher demo publisher for testing and demonistration only
type TimedPublisher struct {
	id               string
	publishToChannel chan (message.Message)
	isRunning        bool
}

// NewTimedPublisher constructor
func NewTimedPublisher(ID string) *TimedPublisher {
	return &TimedPublisher{
		id:               ID,
		publishToChannel: nil,
		isRunning:        false,
	}
}

// ID gets the ID
func (p TimedPublisher) ID() string {
	return p.id
}

// PublishTo publishes to channel
func (p *TimedPublisher) PublishTo(channel chan (message.Message)) {
	p.publishToChannel = channel
}

// Run starts the publisher
func (p *TimedPublisher) Run() error {
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

		p.publishToChannel <- message.Message{
			CreatedAt: time.Now().UTC(),
			ID:        id.String(),
			Payload:   fmt.Sprintf("message from %s", p.id),
			Tags:      nil,
		}
	}

	return nil
}

// Stop stops the publisher
func (p *TimedPublisher) Stop() error {
	p.isRunning = false
	return nil
}
