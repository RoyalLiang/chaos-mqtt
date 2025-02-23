package service

import (
	"errors"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var Redis *redis.Client

func newRedisClient(config *configs.RedisConfig) (*redis.Client, error) {
	if config == nil {
		return nil, errors.New("no redis config provided")
	}
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	}), nil
}

func init() {
	var err error

	Redis, err = newRedisClient(configs.Chaos.Redis)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
