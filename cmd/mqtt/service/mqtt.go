package service

import (
	"context"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"sync"
)

var (
	MQTTClient *mqttClient
)

type mqttClient struct {
	server  mqtt.Client
	address string
	ctx     context.Context
	exit    chan struct{}
}

func (mc *mqttClient) parseMQTTOptions() *mqtt.ClientOptions {
	options := mqtt.NewClientOptions()
	options.ProtocolVersion = 5
	options.SetAutoReconnect(true)
	options.SetClientID("6dd337e1-f0a1-44d8-810f-9485bfde8f4f")
	//options.OnConnect = mc.OnConnect
	if configs.FMSConfig.MQTT.Address != "" {
		mc.address = configs.FMSConfig.MQTT.Address
		options.AddBroker(configs.FMSConfig.MQTT.Address)

		if configs.FMSConfig.MQTT.User != "" {
			options.SetUsername(configs.FMSConfig.MQTT.User)
		}
		if configs.FMSConfig.MQTT.Password != "" {
			options.SetPassword(configs.FMSConfig.MQTT.Password)
		}
		return options
	} else {
		return nil
	}
}

func (mc *mqttClient) Publish(topic string, message interface{}) error {
	if mc.server == nil {
		mc.Init()
	}
	if token := mc.server.Publish(topic, byte(1), false, message); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	mc.server.Disconnect(1000)
	return nil
}

func (mc *mqttClient) OnConnect(c mqtt.Client) {
	fmt.Println("mqtt server connected...")
}

func (mc *mqttClient) Subscribe(wg *sync.WaitGroup, topic string, qos byte, callback mqtt.MessageHandler) {
	defer wg.Done()
	if mc.server == nil {
		mc.Init()
	}
	for {
		if token := mc.server.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
			fmt.Println("subscribe error: ", token.Error())
			os.Exit(1)
		}
	}

}

func (mc *mqttClient) Init() {
	options := MQTTClient.parseMQTTOptions()
	if options == nil {
		fmt.Println("can not found mqtt host...")
		os.Exit(1)
	}
	MQTTClient.server = mqtt.NewClient(options)

	if token := MQTTClient.server.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("can not connect to mqtt server...")
		os.Exit(1)
	}
}

func init() {
	MQTTClient = &mqttClient{
		ctx:  context.Background(),
		exit: make(chan struct{}),
	}
}
