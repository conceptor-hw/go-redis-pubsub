package main

import (
	"encoding/json"
	"log"
	"time"

	"go_pubsub/driver/pubsub"

	"gopkg.in/redis.v2"
)

func main() {
	var pub *redis.IntCmd
	var err error
	var sub_channel string = "rust_channel"
	var pub_channel string = "go_channel"

	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags
	// Create a subscriber
	_, err = pubsub.NewSubscriber(sub_channel, handle_msg)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}

	log.Print("Subscriptions done. Publishing...")
	time.Sleep(time.Second)

	for i := 0; i < 10000; i++ {
		payload := pubsub.Order{Description: "message from go", Quantity: 0, Index: int32(i)}
		// -- Publish some stuf --
		message := pubsub.Message{Id: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			Channel: pub_channel,
			Payload: payload,
		}

		pub = pubsub.Service.Publish(pub_channel, message)
		if err = pub.Err(); err != nil {
			log.Print("PublishString() error", err)
		}
		_ = pub
		time.Sleep(2 * time.Second)
	}

	for {
		time.Sleep(time.Second)
	}
}

func handle_msg(channel, payload string) {
	var msg pubsub.Message
	err := json.Unmarshal([]byte(payload), &msg)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	log.Printf("subcriber msg is: %v ", msg)
}
