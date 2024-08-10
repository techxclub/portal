package client

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/client/azure"
	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/client/email"
	"github.com/techx/portal/client/google"
	"github.com/techx/portal/client/ratelimiter"
	"github.com/techx/portal/config"
)

type Registry struct {
	DB                db.Client
	RateLimiter       ratelimiter.RateLimiter
	GoogleClient      google.Client
	ServiceMailClient email.Client
	SupportMailClient email.Client
	OTPCache          cache.Cache[cache.OTPCache]
	AzureStorage      azure.Client
}

func NewRegistry(cfg *config.Config) *Registry {
	redisClient := newRedisClient(cfg.Redis)
	dbClient := db.NewPostgresDBClient(cfg)
	rateLimiter := ratelimiter.NewRateLimiter(redisClient)
	otpCache := cache.NewOTPCache(redisClient, cfg.Redis)
	googleClient := google.NewGoogleClient(cfg.GoogleClient)
	serviceMailClient := email.NewEmailClient(cfg.ServiceMail)
	supportMailClient := email.NewEmailClient(cfg.SupportMail)
	azureStorage := azure.NewAzureClient(cfg.AzureStorage)

	return &Registry{
		DB:                dbClient,
		RateLimiter:       rateLimiter,
		GoogleClient:      googleClient,
		ServiceMailClient: serviceMailClient,
		SupportMailClient: supportMailClient,
		OTPCache:          otpCache,
		AzureStorage:      azureStorage,
	}
}

func newRedisClient(redisCfg config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:               redisCfg.GetAddress(),
		PoolSize:           redisCfg.GetPoolSize(),
		MinIdleConns:       redisCfg.GetMinIdleConnections(),
		Username:           redisCfg.GetUsername(),
		Password:           redisCfg.GetPassword(),
		DialTimeout:        redisCfg.GetDialTimeout(),
		PoolTimeout:        redisCfg.GetPoolTimeout(),
		ReadTimeout:        redisCfg.GetReadTimeout(),
		WriteTimeout:       redisCfg.GetWriteTimeout(),
		IdleTimeout:        redisCfg.GetIdleTimeout(),
		IdleCheckFrequency: redisCfg.GetIdleCheckFrequency(),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Error().Err(err).Msg("failed to connect to redis")
		panic(err)
	}

	return client
}
