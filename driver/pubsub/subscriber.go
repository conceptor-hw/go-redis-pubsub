package pubsub

import (
	"fmt"
	"log"

	"gopkg.in/redis.v2"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback processFunc
}

type Order struct {
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	Index       int32  `json:"index"`
}

type Message struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
	Payload Order  `json:"payload"`
}

type processFunc func(string, string)

func NewSubscriber(channel string, fn processFunc) (*Subscriber, error) {
	var err error
	// TODO Timeout param?

	s := Subscriber{
		pubsub:   Service.client.PubSub(),
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

	err = s.pubsub.Subscribe(s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, err := s.pubsub.Receive()
		if err != nil {
			fmt.Printf("try subscribe channel[test_channel] error[%s]\n", err.Error())
			continue
		}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		case *redis.PMessage:
			channel = m.Channel
			payload = m.Payload
		}

		// Process the message
		go s.callback(channel, payload)
	}
}
