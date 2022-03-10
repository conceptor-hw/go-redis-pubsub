package main

import (
	"go_pubsub/driver/pubsub"
	"log"
	"time"
)

func main() {
	var err error
	// var sub_prover_channel string = "prover_pub_channel"

	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags
	sub_channel := []string{pubsub.SUB_BINARY_CHANNEL, pubsub.SUB_MGT_CHANNEL, pubsub.SUB_BINARY_CHANNEL_FROM_POOL, pubsub.SUB_MGT_CHANNEL_FROM_POOL}

	// Create a pool server subscriber
	for _, channel_ele := range sub_channel {
		_, err = pubsub.NewSubscriber(channel_ele)
		if err != nil {
			log.Println("NewSubscriber() error", err)
		}
	}

	log.Print("Subscriptions done. Publishing...")
	time.Sleep(time.Second)

	// for i := 0; i < 10000; i++ {

	// 	// -- Publish some stuf --
	// 	payload := pubsub.Order{Description: "message from go", Quantity: 0, Index: int32(i)}
	// 	// -- Publish some stuf --
	// 	message := pubsub.PubSubMessage{Id: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
	// 		Channel: pubsub.PUB_MGT_CHANNEL,
	// 		Payload: payload,
	// 	}
	// 	fmt.Println("publish message ", message)

	// 	err = pubsub.Service.PubNormalMsg(pubsub.PUB_MGT_CHANNEL, message)
	// 	if err != nil {
	// 		log.Print("PublishString() error", err)
	// 	}
	// 	time.Sleep(2 * time.Second)
	// }

	for {

	}
}
