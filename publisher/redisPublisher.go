package publisher

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/husainaloos/event-bus/message"
)

// RedisPublisher is a publisher implementation that reads from redis.
type RedisPublisher struct {
	id               string
	addr             string
	password         string
	redisChannelName string
	client           *redis.Client
	publishTo        chan message.Message
	interrupt        chan bool
	errorStopping    chan error
}

// NewRedisPublisher creates a new RedisPublisher
func NewRedisPublisher(id, addr, password string) *RedisPublisher {
	return &RedisPublisher{
		addr:          addr,
		password:      password,
		id:            id,
		interrupt:     make(chan bool, 1),
		errorStopping: make(chan error, 1),
	}
}

// ID gets the id of the publisher
func (r RedisPublisher) ID() string {
	return r.id
}

// PublishTo sets the channel to which the publisher will publish
func (r *RedisPublisher) PublishTo(c chan message.Message) {
	r.publishTo = c
}

// Run runs the redis publisher
func (r *RedisPublisher) Run() error {
	if r.publishTo == nil {
		return errors.New("publish channel is not set")
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.password,
	})

	sub := r.client.Subscribe(r.redisChannelName)

	go func() {
		for {

			select {
			case <-r.interrupt:
				log.Printf("received interrupt singnal. stopping")
				if err := sub.Close(); err != nil {
					r.errorStopping <- err
				}
				if err := r.client.Close(); err != nil {
					r.errorStopping <- err
				}
				break
			default:
				m, err := sub.ReceiveTimeout(5 * time.Second)
				if err != nil {
					log.Printf("cannot read from the redis channel: %+v", err)
				} else {
					log.Printf("%v", m)
				}
			}
		}
	}()

	return nil
}

// Stop stops the redis publisher
func (r *RedisPublisher) Stop() error {
	r.interrupt <- true
	err := <-r.errorStopping
	return err
}
