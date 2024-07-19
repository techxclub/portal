package client

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/client/ratelimiter"
	"github.com/techx/portal/config"
	"gopkg.in/gomail.v2"
)

type Registry struct {
	DB                 db.Client
	RateLimiter        ratelimiter.RateLimiter
	ReferralMailClient *gomail.Dialer
	OTPMailClient      *gomail.Dialer
	OTPCache           cache.Cache[cache.OTPCache]
}

func NewRegistry(cfg *config.Config) *Registry {
	redisClient := newRedisClient(cfg.Redis)
	dbClient := db.NewPostgresDBClient(cfg)
	rateLimiter := ratelimiter.NewRateLimiter(redisClient)
	otpCache := cache.NewOTPCache(redisClient, cfg.Redis)
	referralMailClient := newGmailClient(cfg.ReferralMail)
	otpMailClient := newGmailClient(cfg.OTPMail)

	return &Registry{
		DB:                 dbClient,
		RateLimiter:        rateLimiter,
		ReferralMailClient: referralMailClient,
		OTPMailClient:      otpMailClient,
		OTPCache:           otpCache,
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

func newGmailClient(gmailCfg config.MailSMTP) *gomail.Dialer {
	return gomail.NewDialer(
		gmailCfg.SMTPServer,
		gmailCfg.SMTPPort,
		gmailCfg.FromEmail,
		gmailCfg.SMTPPassword,
	)
}
