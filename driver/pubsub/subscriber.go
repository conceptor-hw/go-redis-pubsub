package pubsub

import (
	"context"
	"encoding/json"
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
			handle_prover_spec_message(channel, payload)
		default:
			panic("unreached")
		}

	}
}

type ProveSpecMessage struct {
	Prover_id string
	Info      string
}

func PubProSpecMsg(channel string, data interface{}) error {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	fmt.Println("messagei is", messageString)
	err = R_client.Publish(R_ctx, channel, messageString).Err()
	if err != nil {
		return errors.New("publish normal message wrong... " + err.Error())
	}
	return nil
}

func handle_prover_spec_message(channel string, payload string) {
	var pub_channel string = "binary_channel_schedule"
	pub_info := ProveSpecMessage{Prover_id: "abc", Info: payload}
	err := PubProSpecMsg(pub_channel, pub_info)
	if err != nil {
		log.Print("PublishString() error", err)
	}
}
