package pubsub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const (
	SUB_BINARY_CHANNEL           = "binary_channel_prover"            // 订阅来自prover的二进制message的channel
	PUB_BINARY_CHANNEL           = "binary_channel_schedule"          // 向prover发布二进制meesage的channel
	SUB_MGT_CHANNEL              = "mgt_channel_prover"               // 订阅来自prover的控制面的message的channel
	PUB_MGT_CHANNEL              = "mgt_channel_schedule"             // 向prover发布控制面的message的channel
	SUB_BINARY_CHANNEL_FROM_POOL = "binary_channel_pool"              // 订阅来自pool server的二进制message的channel
	PUB_BINARY_CHANNEL_FOR_POOL  = "binary_channel_schedule_for_pool" // 向prover发布二进制meesage的channel
	SUB_MGT_CHANNEL_FROM_POOL    = "mgt_channel_pool"                 // 订阅来自pool server的控制面的message channel
	PUB_MGT_CHANNEL_FOR_POOL     = "mgt_channel_schedule_for_pool"    // 向pool server发布控制面的message的channel
)

type PubSub struct {
	client *redis.Client
	ctx    context.Context
}

var Service *PubSub
var pub *redis.IntCmd = nil

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	Service = &PubSub{client, ctx}
}

func (ps *PubSub) PubBinaryMsg(channel string, data []byte) error {

	fmt.Println("binary channel", channel)
	err := ps.client.Publish(ps.ctx, channel, data).Err()
	if err != nil {
		return errors.New("publish Data wrong... " + err.Error())
	}
	return nil
}

func (ps *PubSub) PubNormalMsg(channel string, data interface{}) error {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	err = ps.client.Publish(ps.ctx, channel, messageString).Err()
	if err != nil {
		return errors.New("publish normal message wrong... " + err.Error())
	}
	return nil
}
