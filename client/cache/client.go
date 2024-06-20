package cache

import (
	"context"
	"time"

	"github.com/techx/portal/errors"
)

type Cache[T any] interface {
	Set(ctx context.Context, key string, value T, ttl time.Duration) error
	Get(ctx context.Context, key string) (T, error)
	Del(ctx context.Context, key string) error
}

type otpCache struct {
	store map[string]string
}

func NewCache() Cache[string] {
	return otpCache{
		store: make(map[string]string),
	}
}

func (c otpCache) Set(_ context.Context, key, value string, _ time.Duration) error {
	if value == "" {
		return errors.ErrValueCannotBeEmpty
	}

	if key == "" {
		return errors.ErrKeyCannotBeEmpty
	}

	c.store[key] = value
	return nil
}

func (c otpCache) Get(_ context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.ErrKeyCannotBeEmpty
	}

	value, ok := c.store[key]
	if !ok {
		return "", errors.ErrKeyNotFound
	}

	return value, nil
}

func (c otpCache) Del(_ context.Context, key string) error {
	if key == "" {
		return errors.ErrKeyCannotBeEmpty
	}

	delete(c.store, key)
	return nil
}
