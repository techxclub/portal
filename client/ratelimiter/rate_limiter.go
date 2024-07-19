package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RateLimiter represents interface to instantiate a rate limiter.
type RateLimiter interface {
	TryAcquire(ctx context.Context, key string, attempts, windowSecs int64) (bool, error)
	Reset(ctx context.Context, key string) error
}

type RateLimitConfig struct {
	Attempts   int64
	WindowSecs int64
}

type rateLimiter struct {
	redisClient *redis.Client
}

func NewRateLimiter(redisClient *redis.Client) RateLimiter {
	return &rateLimiter{
		redisClient: redisClient,
	}
}

func (rl *rateLimiter) TryAcquire(ctx context.Context, key string, attempts, windowSecs int64) (bool, error) {
	return rl.slidingWindowRateLimitWithAttempts(ctx, key, attempts, windowSecs)
}

func (rl *rateLimiter) Reset(ctx context.Context, key string) error {
	return rl.redisClient.Del(ctx, key).Err()
}

func (rl *rateLimiter) slidingWindowRateLimitWithAttempts(ctx context.Context, key string, attempts, windowSecs int64) (bool, error) {
	now := currentTimeInNano()
	windowTimeInNanoSecs := toNanoSeconds(windowSecs)

	// Start a pipeline to execute a batch of commands atomically.
	pipe := rl.redisClient.Pipeline()

	// Remove scores (members) outside the current window.
	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprint(now-windowTimeInNanoSecs))
	// Count the number of members (attempts) in the current window.
	cardCmd := pipe.ZCard(ctx, key)
	// Add the current attempt with the current timestamp as score.
	pipe.ZAdd(ctx, key, &redis.Z{Score: float64(now), Member: now})
	// Set the expiration of the key to the length of the window.
	pipe.Expire(ctx, key, time.Duration(windowSecs)*time.Second)

	// Execute the pipeline.
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// Get the result of the ZCard command to find out the number of attempts in the window.
	data, err := cardCmd.Result()
	if err != nil {
		return false, err
	}

	// If the number of attempts is less than the allowed attempts, the action is allowed.
	return data < attempts, nil
}

func currentTimeInNano() int64 {
	return time.Now().UnixNano()
}

func toNanoSeconds(seconds int64) int64 {
	return seconds * 1e9
}
