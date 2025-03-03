package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"

	"fms-awesome-tools/configs"
)

func TestRedis(t *testing.T) {
	cfg := &configs.RedisConfig{
		Address:  "10.1.205.3:16397",
		DB:       0,
		Password: "aeiou",
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", cfg.Address),
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
