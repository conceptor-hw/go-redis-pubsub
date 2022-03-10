package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Subscriber struct {
	pubsub  *redis.PubSub
	ctx     context.Context
	channel string
}

type processFunc func(string, string)

type Order struct {
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	Index       int32  `json:"index"`
}

type PubSubMessage struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
	Payload Order  `json:"payload"`
}

func NewSubscriber(channel string) (*Subscriber, error) {
	var err error
	// TODO Timeout param?

	s := Subscriber{
		pubsub:  nil,
		ctx:     context.Background(),
		channel: channel,
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

	s.pubsub = Service.client.Subscribe(s.ctx, s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.", s.channel)
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, err := s.pubsub.Receive(s.ctx)
		if err != nil {
			fmt.Printf("try subscribe channel[test_channel] error[%s]\n", err.Error())
			continue
		}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", m.Channel)

		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
			handle_msg(channel, payload)
		default:
			panic("unreached")
		}

	}
}

func handle_msg(channel, payload string) {
	fmt.Println("handle message from: ", channel)
	switch channel {
	case SUB_BINARY_CHANNEL:
		fmt.Println("#####SUB_BINARY_CHANNEL11111: ")
		Service.PubBinaryMsg(PUB_BINARY_CHANNEL_FOR_POOL, []byte(payload))

	case SUB_MGT_CHANNEL, SUB_MGT_CHANNEL_FROM_POOL:
		{
			var msg PubSubMessage
			err := json.Unmarshal([]byte(payload), &msg)
			if err != nil {
				log.Printf("Unmarshal error: %v", err)
			}
			log.Printf("subcriber msg is: %v ", msg)
		}
	case SUB_BINARY_CHANNEL_FROM_POOL:
		fmt.Println("#####SUB_BINARY_CHANNEL_FROM_POOL222222: ")
		Service.PubBinaryMsg(PUB_BINARY_CHANNEL, []byte(payload))
	default:
		fmt.Println("sub channel is wrong ....", channel)
	}
}
