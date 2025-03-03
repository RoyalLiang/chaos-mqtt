package service

import (
	"errors"

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
