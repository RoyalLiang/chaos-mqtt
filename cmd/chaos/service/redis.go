package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"fms-awesome-tools/configs"
)

func NewRedisClient() (*redis.Client, error) {

	if configs.Chaos.Redis == nil {
		return nil, errors.New("no redis config provided")
	}
	return redis.NewClient(&redis.Options{
		Addr:     configs.Chaos.Redis.Address,
		Password: configs.Chaos.Redis.Password,
		DB:       configs.Chaos.Redis.DB,
	}), nil
}

func Subscribe(ctx context.Context, channel string, msgChan chan *redis.Message) error {
	rc, err := NewRedisClient()
	if err != nil {
		return err
	}

	sub := rc.Subscribe(ctx, channel)
	defer sub.Close()

	go func() {
		for {
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				fmt.Println("subs error:", err.Error())
				close(msgChan)
				return
			}
			msgChan <- msg
		}
	}()

	return nil
}
