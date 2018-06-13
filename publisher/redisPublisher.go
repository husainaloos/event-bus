package publisher

import (
	"errors"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/husainaloos/event-bus/message"
)

// RedisPublisher is a publisher implementation that reads from redis.
type RedisPublisher struct {
	id               string
	addr             string
	password         string
	redisChannelName string
	publishTo        chan message.Message
	redisConn        redis.Conn
	closing          chan bool
	closed           chan bool
}

// NewRedisPublisher creates a new RedisPublisher
func NewRedisPublisher(id, addr, password, channelName string) *RedisPublisher {
	return &RedisPublisher{
		addr:             addr,
		password:         password,
		id:               id,
		redisChannelName: channelName,
		redisConn:        nil,
		closing:          make(chan bool, 1),
		closed:           make(chan bool, 1),
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

	var err error
	r.redisConn, err = redis.Dial("tcp", r.addr, redis.DialKeepAlive(1*time.Minute), redis.DialPassword(r.password))
	if err != nil {
		return err
	}

	pubSubConn := redis.PubSubConn{
		Conn: r.redisConn,
	}

	if err := pubSubConn.PSubscribe(r.redisChannelName); err != nil {
		return err
	}

	go func(pubSubConn redis.PubSubConn) {
		for {
			switch response := pubSubConn.Receive().(type) {
			case redis.PMessage:
				log.Printf("%s received message from channel %s: %+v", r.ID(), response.Channel, response)
				r.publishTo <- message.Message{
					ID:        "someId",
					CreatedAt: time.Now().UTC(),
					Tags:      nil,
					Payload:   string(response.Data),
				}
			case redis.Subscription:
				log.Printf("%s: received subscription message: %+v", r.ID(), response)
			case redis.Pong:
				log.Printf("%s: received pong: %+v", r.ID(), response)
			case error:
				select {
				case <-r.closing:
					log.Printf("%s: closing redis publisher.", r.ID())
					r.closed <- true
					break
				default:
					log.Fatalf("%s: the message received from redis is unrecognized. %+v", r.ID(), response)
				}
			default:
				log.Printf("%s: not sure what's going on. %T %+v", r.ID(), response, response)
			}
		}
	}(pubSubConn)

	return nil
}

// Stop stops the redis publisher
func (r *RedisPublisher) Stop() error {
	log.Printf("%s: closing", r.ID())
	r.closing <- true
	err := r.redisConn.Close()
	<-r.closed
	return err
}
