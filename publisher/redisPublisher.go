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
}

// NewRedisPublisher creates a new RedisPublisher
func NewRedisPublisher(id, addr, password, channelName string) *RedisPublisher {
	return &RedisPublisher{
		addr:             addr,
		password:         password,
		id:               id,
		redisChannelName: channelName,
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

	conn, err := redis.Dial("tcp", r.addr, redis.DialKeepAlive(1*time.Minute), redis.DialPassword(r.password))
	if err != nil {
		log.Fatalf("%s: unable to connect to redis. %+v", r.ID(), err)
	}

	pubSubConn := redis.PubSubConn{
		Conn: conn,
	}

	pubSubConn.PSubscribe(r.redisChannelName)
	go func(pubSubConn redis.PubSubConn) {
		for {
			response := pubSubConn.ReceiveWithTimeout(1 * time.Second)

			sub, ok := response.(redis.Subscription)
			if ok {
				log.Printf("%s: received subscription: %+v", r.ID(), sub)
				continue
			}

			mes, ok := response.(redis.PMessage)
			if ok {
				log.Printf("%s received message from channel %s: %+v", r.ID(), mes.Channel, mes)

				r.publishTo <- message.Message{
					ID:        "someId",
					CreatedAt: time.Now().UTC(),
					Tags:      nil,
					Payload:   string(mes.Data),
				}

				continue
			}

			pong, ok := response.(redis.Pong)
			if ok {
				log.Printf("%s: received pong: %+v", r.ID(), pong)
				continue
			}

			err, ok := response.(redis.Error)
			if ok {
				log.Fatalf("%s: received err: %+v", r.ID(), err)
				continue
			}

			log.Fatalf("%s: the message received from redis is unrecognized. %+v", r.ID(), response)
		}
	}(pubSubConn)

	return nil
}

// Stop stops the redis publisher
func (r *RedisPublisher) Stop() error {
	return nil
}
