package redis

import (
	"awesomeProject/internal/app/config"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const servicePrefix = "container_logistics."

type Client struct {
	cfg    config.RedisConfig
	client *redis.Client
}

func New(cfg config.RedisConfig) (*Client, error) {
	client := &Client{}
	client.cfg = cfg

	redisClient := redis.NewClient(&redis.Options{
		Password:    cfg.Password,
		Username:    cfg.User,
		Addr:        cfg.Host + ":" + strconv.Itoa(cfg.Port),
		DB:          0,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	})

	client.client = redisClient

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("cant ping redis: %w", err)
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
