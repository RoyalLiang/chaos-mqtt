package service

import (
	"context"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestRedis(t *testing.T) {
	cfg := &configs.RedisConfig{
		Host:     "10.1.205.3",
		Port:     16397,
		DB:       0,
		Password: "aeiou",
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	var ctx = context.Background()

	sub := client.Subscribe(ctx, "vehicle_status")

	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
			t.Fatal()
		}
		fmt.Println("received", msg.Payload, "from", msg.Channel)
	}

}
