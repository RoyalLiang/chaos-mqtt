package service

import (
	"context"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"

	tools "fms-awesome-tools/utils"
)

var sub subscriber

type subscriber struct {
	exit        struct{}
	client      *MqttClient
	redisClient *redis.Client
}

func (s subscriber) msgHandler(client mqtt.Client, message mqtt.Message) {
	topic := message.Topic()
	msg := string(message.Payload())
	now := time.Now().Format("2006-01-02 15:04:051")
	fmt.Printf("[%s] receive msg from %s ==> %s \n", now, topic, msg)
}

func StartSubscribe(topic string) {
	fmt.Println(tools.CustomTitle("\n          Chaos Subscribe Start Listen...          \n"))
	sub.client.Subscribe(topic, 0, sub.msgHandler)
}

func StartRedisSubscribe(topic string) {
	fmt.Println(tools.CustomTitle("\n          Chaos Redis Subscribe Start To Listen...          \n"))

	cli, err := NewRedisClient()
	if err != nil {
		cobra.CheckErr(err)
	}

	ctx := context.Background()
	subs := cli.Subscribe(ctx, topic)
	defer subs.Close()
	for {
		msg, err := subs.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("subs error:", err.Error())
			return
		}
		fmt.Println("receive msg: ", msg)
	}
}

func init() {
	var err error
	sub = subscriber{}

	sub.client, err = NewMQTTClientWithConfig("subs")
	if err != nil {
		fmt.Println("error for init mqtt client: ", err)
		os.Exit(1)
	}
}
