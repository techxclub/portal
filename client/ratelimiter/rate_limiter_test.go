package ratelimiter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/techx/portal/client/ratelimiter"
)

type rateLimiterTest struct {
	suite.Suite
	ctx         context.Context
	redisClient *redis.Client
}

func TestRateLimiterTest(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(rateLimiterTest))
}

func (s *rateLimiterTest) SetupSuite() {
	s.ctx = context.Background()
	s.redisClient = redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		PoolSize:           10,
		MinIdleConns:       5,
		DialTimeout:        5 * time.Second,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       5 * time.Second,
		IdleTimeout:        5 * time.Second,
		IdleCheckFrequency: 5 * time.Second,
	})

	if err := s.redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func (s *rateLimiterTest) TearDownSuite() {
	if err := s.redisClient.Close(); err != nil {
		panic(err)
	}
}

func (s *rateLimiterTest) TestRateLimiter_TryAcquire() {
	attempts := int64(2)
	windowSecs := int64(2)
	limiter := ratelimiter.NewRateLimiter(s.redisClient)

	s.Run("Should acquire if rate limit not exceeded", func() {
		key := uniqueKey("xyz")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt2)
	})

	s.Run("Should not acquire if rate limit exceeded", func() {
		key := uniqueKey("xxx")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt2)

		attempt3, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.False(attempt3)
	})

	s.Run("Should acquire if window exhausted", func() {
		key := uniqueKey("ooo")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt2)

		time.Sleep(5 * time.Second)

		attempt3, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt3)
	})
}

func (s *rateLimiterTest) TestRateLimiter_TryAcquireWithOverriddenConfig() {
	attempts := int64(2)
	windowSecs := int64(2)
	limiter := ratelimiter.NewRateLimiter(s.redisClient)

	s.Run("Should acquire if rate limit not exceeded", func() {
		key := uniqueKey("xyz")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt2)
	})

	s.Run("Should not acquire if rate limit exceeded", func() {
		key := uniqueKey("xxxx")
		attempt1, err := limiter.TryAcquire(s.ctx, key, 1, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, 1, windowSecs)
		require.NoError(s.T(), err)
		s.False(attempt2)
	})

	s.Run("Should acquire if window exhausted", func() {
		key := uniqueKey("oooo")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, 3)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, 3)
		require.NoError(s.T(), err)
		s.True(attempt2)

		time.Sleep(5 * time.Second)

		attempt3, err := limiter.TryAcquire(s.ctx, key, 1, 3)
		require.NoError(s.T(), err)
		s.True(attempt3)

		attempt4, err := limiter.TryAcquire(s.ctx, key, 1, 3)
		require.NoError(s.T(), err)
		s.False(attempt4)
	})
}

func (s *rateLimiterTest) TestRateLimiter_Reset() {
	attempts := int64(1)
	windowSecs := int64(10)
	limiter := ratelimiter.NewRateLimiter(s.redisClient)

	s.Run("Can't acquire if rate limit exceeded", func() {
		key := uniqueKey("xxx")
		attempt1, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt1)

		attempt2, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.False(attempt2)

		err = limiter.Reset(s.ctx, key)
		require.NoError(s.T(), err)

		attempt3, err := limiter.TryAcquire(s.ctx, key, attempts, windowSecs)
		require.NoError(s.T(), err)
		s.True(attempt3)
	})
}

func uniqueKey(base string) string {
	return fmt.Sprintf("%s_%d", base, time.Now().UnixNano())
}
