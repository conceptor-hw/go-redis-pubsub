package main

import (
	"go_pubsub/driver/pubsub"
	"log"
	"time"
)

func main() {
	var err error
	var sub_channel string = "binary_channel_prover"

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
		// message := pubsub.ProverMessage{Previous_block_hash: "191cbbfc488440ce95e9d5d0770d8c65",
		// 	Block_height:      uint32(0),
		// 	Block_timestamp:   0,
		// 	Difficulty_target: uint64(0),
		// }
		// temp_str := fmt.Sprint(message)
		// pub_info := pubsub.ProveSpecMessage{Prover_id: "abc", Info: temp_str}
		// err = pubsub.PubNormalMsg(pub_channel, pub_info)
		// if err != nil {
		// 	log.Print("PublishString() error", err)
		// }
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
