package publisher

import (
	"fmt"
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
	pubSub           *redis.PubSubConn
}

// NewRedisPublisher creates a new RedisPublisher
func NewRedisPublisher(id, addr, password, channelName string) *RedisPublisher {
	return &RedisPublisher{
		addr:             addr,
		password:         password,
		id:               id,
		redisChannelName: channelName,
		redisConn:        nil,
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

	r.pubSub = &redis.PubSubConn{
		Conn: r.redisConn,
	}

	if err := r.pubSub.PSubscribe(r.redisChannelName); err != nil {
		return err
	}

	go func(pubSubConn *redis.PubSubConn) {
		for {
			switch response := r.pubSub.Receive().(type) {
			case redis.PMessage:
				log.Printf("%s received message from channel %s: %+v", r.ID(), response.Channel, response)
				r.publishTo <- message.Message{
					ID:        "someId",
					CreatedAt: time.Now().UTC(),
					Tags:      nil,
					Payload:   string(response.Data),
				}
			case redis.Subscription:
				log.Printf("%s: %s to %s", r.ID(), response.Kind, response.Channel)
			case redis.Pong:
				log.Printf("%s: received pong: %+v", r.ID(), response)
			case error:
				log.Fatalf("%s: the message received from redis is unrecognized. %+v", r.ID(), response)
			default:
				log.Printf("%s: not sure what's going on. %T %+v", r.ID(), response, response)
			}
		}
	}(r.pubSub)

	return nil
}

// Stop stops the redis publisher
func (r *RedisPublisher) Stop() error {
	log.Printf("%s: closing", r.ID())

	if err := r.pubSub.PUnsubscribe("*"); err != nil {
		return fmt.Errorf("cannot close redis subscription: %+v", err);
	}

	if err := r.redisConn.Close(); err != nil{
		return fmt.Errorf("cannot close connection to redis: %+v", err);
	}

	return nil
}
