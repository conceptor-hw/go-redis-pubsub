package pubsub

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	ctx      context.Context
	channel  string
	callback processFunc
}

type processFunc func(string, string)

func NewSubscriber(channel string, fn processFunc) (*Subscriber, error) {
	var err error
	// TODO Timeout param?

	s := Subscriber{
		pubsub:   nil,
		ctx:      context.Background(),
		channel:  channel,
		callback: fn,
	}

	// Subscribe to the channel
	err = s.subscribe()
	if err != nil {
		return nil, err
	}

	// Listen for messages
	log.Printf("begin to goroute listen....")
	go s.listen()

	return &s, nil
}

func (s *Subscriber) subscribe() error {
	var err error

	s.pubsub = R_client.Subscribe(s.ctx, s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	// var channel string
	// var payload string

	for {
		msg, err := s.pubsub.Receive(s.ctx)
		if err != nil {
			fmt.Printf("try subscribe channel[test_channel] error[%s]\n", err.Error())
			continue
		}

		// channel = ""
		// payload = ""

		fmt.Println("recv msg ####", msg)
		switch msg := msg.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", msg.Channel)

		case *redis.Message:
			fmt.Println("received!!!!", msg.Payload, "from", msg.Channel, "payloadSlice", msg.PayloadSlice)
			// buf := new(bytes.Buffer)
			// b := buf.Bytes()
			// dec := gob.NewDecoder(bytes.NewBuffer(b))
			// err = dec.Decode(&msg.PayloadSlice)

		default:
			panic("unreached")
		}

	}
}
