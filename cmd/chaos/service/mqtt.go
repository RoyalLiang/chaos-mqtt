package service

import (
	"context"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/google/uuid"
	"os"

	"github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client  mqtt.Client
	address string
	ctx     context.Context
	exit    chan struct{}
}

func (mc *MqttClient) connectToServer(prefix, user, password string) error {
	options := mqtt.NewClientOptions()
	options.ProtocolVersion = 5
	options.SetAutoReconnect(true)
	options.SetClientID(configs.Chaos.Product.Name + "-" + prefix + "-" + uuid.NewString())
	options.AddBroker(mc.address)
	options.SetUsername(user)
	options.SetPassword(password)

	mc.client = mqtt.NewClient(options)
	if token := mc.client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("mqtt connect error: %s\n", token.Error())
	}
	return nil
}

func NewMQTTClient(prefix, address, user, password string) (*MqttClient, error) {
	if address == "" {
		return nil, fmt.Errorf("address is empty")
	}

	c := &MqttClient{
		address: address,
		ctx:     context.Background(),
		exit:    make(chan struct{}),
	}
	if err := c.connectToServer(prefix, user, password); err != nil {
		return nil, err
	}
	return c, nil
}

func NewMQTTClientWithConfig(prefix string) (*MqttClient, error) {
	config := configs.Chaos.MQTT
	return NewMQTTClient(prefix, config.Address, config.User, config.Password)
}

func (mc *MqttClient) Close() {
	close(mc.exit)
	defer mc.client.Disconnect(500)
}

func (mc *MqttClient) Publish(topic string, message interface{}) error {
	if token := mc.client.Publish(topic, byte(1), false, message); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (mc *MqttClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) {
	if token := mc.client.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		fmt.Println("subscribe error: ", token.Error())
		os.Exit(1)
	}

	//select {
	//case <-mc.ctx.Done():
	//	fmt.Println("mqtt server disconnected...")
	//	return
	//case <-mc.exit:
	//	fmt.Println("mqtt client closed...")
	//}
	<-mc.exit
}

func (mc *MqttClient) SubscribeMultiple(topics map[string]byte, callback mqtt.MessageHandler) {
	if token := mc.client.SubscribeMultiple(topics, callback); token.Wait() && token.Error() != nil {
		fmt.Println("multiple subscribe error: ", token.Error())
		os.Exit(1)
	}

	<-mc.exit
	//select {
	//case <-mc.ctx.Done():
	//	mc.server.Disconnect(200)
	//	fmt.Println("mqtt server disconnected...")
	//	return
	//case <-mc.exit:
	//	mc.server.Disconnect(200)
	//	fmt.Println("mqtt client closed...")
	//}
}
