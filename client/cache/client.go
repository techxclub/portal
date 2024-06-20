package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-redis/redis/v8"
	"github.com/techx/portal/config"
)

const (
	Set = "set"
	Get = "get"
	Del = "del"
)

type Cache[T any] interface {
	Set(ctx context.Context, key string, value *T, ttl time.Duration) error
	Get(ctx context.Context, key string) (*T, error)
	Del(ctx context.Context, key string) error
}

type cacheClient[T any] struct {
	commandPrefix string
	redisClient   redis.Cmdable
}

func NewCache[T any](redisClient redis.Cmdable, commandPrefix string, cfg config.Redis) Cache[T] {
	c := cacheClient[T]{
		commandPrefix: commandPrefix,
		redisClient:   redisClient,
	}

	hystrix.ConfigureCommand(c.getCommandName(Set), cfg.HystrixConfig())
	hystrix.ConfigureCommand(c.getCommandName(Get), cfg.HystrixConfig())
	hystrix.ConfigureCommand(c.getCommandName(Del), cfg.HystrixConfig())
	return &c
}

func (c *cacheClient[T]) Set(ctx context.Context, key string, value *T, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = hystrix.DoC(ctx, c.getCommandName(Set), func(ctx context.Context) error {
		err = c.redisClient.Set(ctx, c.getNamespacedKey(key), data, ttl).Err()
		return err
	}, nil)

	return err
}

func (c *cacheClient[T]) Get(ctx context.Context, key string) (*T, error) {
	var resp string
	var cacheErr error
	err := hystrix.DoC(ctx, c.getCommandName(Get), func(ctx context.Context) error {
		resp, cacheErr = c.redisClient.Get(ctx, c.getNamespacedKey(key)).Result()
		if errors.Is(redis.Nil, cacheErr) {
			return nil
		}
		return cacheErr
	}, nil)
	if err != nil {
		return nil, err
	}
	if cacheErr != nil {
		return nil, cacheErr
	}

	var cacheData T
	if err := json.Unmarshal([]byte(resp), &cacheData); err != nil {
		return nil, err
	}

	return &cacheData, nil
}

func (c *cacheClient[T]) Del(ctx context.Context, key string) error {
	return hystrix.DoC(ctx, c.getCommandName(Del), func(ctx context.Context) error {
		err := c.redisClient.Del(ctx, c.getNamespacedKey(key)).Err()
		return err
	}, nil)
}

func (c *cacheClient[T]) getCommandName(op string) string {
	if len(c.commandPrefix) == 0 {
		return op
	}

	return fmt.Sprintf("%s_%s", c.commandPrefix, op)
}

func (c *cacheClient[T]) getNamespacedKey(key string) string {
	if len(c.commandPrefix) == 0 {
		return key
	}

	return fmt.Sprintf("%s:%s", c.commandPrefix, key)
}
