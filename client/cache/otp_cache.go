package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/techx/portal/config"
)

const (
	OTPCachePrefix = "otp_cache"
)

type OTPCache struct {
	OTP        string    `json:"otp"`
	ExpiryTime time.Time `json:"expiry_time"`
	Attempts   int       `json:"attempts"`
}

func NewOTPCache(redisClient redis.Cmdable, cfg config.Redis) Cache[OTPCache] {
	return NewCache[OTPCache](redisClient, OTPCachePrefix, cfg)
}
