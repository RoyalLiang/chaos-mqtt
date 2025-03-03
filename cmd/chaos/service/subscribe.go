package service

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	tools "fms-awesome-tools/utils"
)

var sub subscriber

type subscriber struct {
	exit   struct{}
	client *MqttClient
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

func init() {
	var err error
	sub = subscriber{}

	sub.client, err = NewMQTTClientWithConfig("subs")
	if err != nil {
		fmt.Println("error for init mqtt client: ", err)
		os.Exit(1)
	}
}
