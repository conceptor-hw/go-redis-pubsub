package main

import (
	"fmt"
	"go_pubsub/driver/pubsub"
	"log"
	"time"
)

func main() {
	var err error
	var sub_channel string = "rust_channel"
	var pub_channel string = "go_channel"
	// var sub_prover_channel string = "prover_pub_channel"

	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags
	// Create a pool server subscriber
	_, err = pubsub.NewSubscriber(sub_channel, handle_msg)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}

	// Create a prover server subscriber
	// _, err = pubsub.NewSubscriber(sub_prover_channel, handle_msg)
	// if err != nil {
	// 	log.Println("NewSubscriber() error", err)
	// }

	log.Print("Subscriptions done. Publishing...")
	time.Sleep(time.Second)

	for i := 0; i < 10000; i++ {

		// -- Publish some stuf --
		message := pubsub.ProverMessage{Previous_block_hash: "191cbbfc488440ce95e9d5d0770d8c65",
			Block_height:      uint32(0),
			Block_timestamp:   0,
			Difficulty_target: uint64(0),
		}
		fmt.Println("publish message ", message)

		err = pubsub.PubBinaryData(pub_channel, message)
		if err != nil {
			log.Print("PublishString() error", err)
		}
		time.Sleep(2 * time.Second)
	}

}

func handle_msg(channel, payload string) {
	// var msg pubsub.Message
	// err := json.Unmarshal([]byte(payload), &msg)
	// if err != nil {
	// 	log.Printf("Unmarshal error: %v", err)
	// }

	// log.Printf("subcriber msg is: %v ", msg)
}
