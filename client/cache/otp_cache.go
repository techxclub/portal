package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/techx/portal/config"
)

const (
	OTPCachePrefix = "otp_cache"
)

type OTPCache struct {
	OTP      string `json:"otp"`
	Attempts int    `json:"attempts"`
	Verified bool   `json:"verified"`
}

func NewOTPCache(redisClient redis.Cmdable, cfg config.Redis) Cache[OTPCache] {
	return NewCache[OTPCache](redisClient, OTPCachePrefix, cfg)
}
