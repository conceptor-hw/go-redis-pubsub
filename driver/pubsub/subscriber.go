package pubsub

import (
	"context"
	"errors"
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

type ProveSpecMessage struct {
	Prover_id string
	Info      string
}

func handle_msg(channel, payload string) error {
	fmt.Println("handle message from: ", channel)
	var pub_channel string = "binary_channel_schedule"
	switch channel {
	case "binary_channel_prover":
		fmt.Println("#####prove#########: ", payload)
		data := ProveSpecMessage{"ab+1", payload}
		temp_str := fmt.Sprint(data)
		binnay_dat := []byte(temp_str)
		err := R_client.Publish(R_ctx, pub_channel, binnay_dat).Err()
		if err != nil {
			return errors.New("publish Data wrong... " + err.Error())
		}
	default:
		fmt.Println("sub channel is wrong ....", channel)
	}
	return nil
}
