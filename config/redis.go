package config

import (
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

type Redis struct {
	Host                   string        `yaml:"HOST" env:"HOST"`
	Port                   int           `yaml:"PORT" env:"PORT"`
	PoolSize               int           `yaml:"POOL_SIZE" env:"POOL_SIZE"`
	MinIdleConnections     int           `yaml:"MIN_IDLE_CONNECTIONS" env:"MIN_IDLE_CONNECTIONS"`
	DialTimeout            time.Duration `yaml:"DIAL_TIMEOUT" env:"DIAL_TIMEOUT"`
	PoolTimeout            time.Duration `yaml:"POOL_TIMEOUT" env:"POOL_TIMEOUT"`
	ReadTimeout            time.Duration `yaml:"READ_TIMEOUT" env:"READ_TIMEOUT"`
	WriteTimeout           time.Duration `yaml:"WRITE_TIMEOUT" env:"WRITE_TIMEOUT"`
	IdleTimeout            time.Duration `yaml:"IDLE_TIMEOUT" env:"IDLE_TIMEOUT"`
	IdleCheckFrequency     time.Duration `yaml:"IDLE_CHECK_FREQUENCY" env:"IDLE_CHECK_FREQUENCY"`
	Username               string        `yaml:"USERNAME" env:"USERNAME"`
	Password               string        `yaml:"PASSWORD" env:"PASSWORD"`
	HystrixTimeout         int           `yaml:"HYSTRIX_TIMEOUT" env:"HYSTRIX_TIMEOUT"`
	MaxConcurrentRequests  int           `yaml:"MAX_CONCURRENT_REQUESTS" env:"MAX_CONCURRENT_REQUESTS"`
	SleepWindow            int           `yaml:"SLEEP_WINDOW" env:"SLEEP_WINDOW"`
	ErrorPercentThreshold  int           `yaml:"ERROR_PERCENT_THRESHOLD" env:"ERROR_PERCENT_THRESHOLD"`
	RequestVolumeThreshold int           `yaml:"REQUEST_VOLUME_THRESHOLD" env:"REQUEST_VOLUME_THRESHOLD"`
}

func defaultRedisConfig() Redis {
	return Redis{
		Host:               "localhost",
		Port:               6379,
		PoolSize:           30,
		MinIdleConnections: 10,
		DialTimeout:        1000 * time.Millisecond,
		PoolTimeout:        1000 * time.Millisecond,
		ReadTimeout:        1000 * time.Millisecond,
		WriteTimeout:       1000 * time.Millisecond,
		IdleTimeout:        30 * time.Minute,
		IdleCheckFrequency: 5 * time.Minute,

		HystrixTimeout:         1000,
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 100,
		SleepWindow:            100,
		ErrorPercentThreshold:  10,
	}
}

func (c *Redis) HystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                c.HystrixTimeout,
		MaxConcurrentRequests:  c.MaxConcurrentRequests,
		SleepWindow:            c.SleepWindow,
		ErrorPercentThreshold:  c.ErrorPercentThreshold,
		RequestVolumeThreshold: c.RequestVolumeThreshold,
	}
}

func (c *Redis) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Redis) GetPoolSize() int {
	return c.PoolSize
}

func (c *Redis) GetMinIdleConnections() int {
	return c.MinIdleConnections
}

func (c *Redis) GetUsername() string {
	return c.Username
}

func (c *Redis) GetPassword() string {
	return c.Password
}

func (c *Redis) GetDialTimeout() time.Duration {
	return c.DialTimeout
}

func (c *Redis) GetPoolTimeout() time.Duration {
	return c.PoolTimeout
}

func (c *Redis) GetReadTimeout() time.Duration {
	return c.ReadTimeout
}

func (c *Redis) GetWriteTimeout() time.Duration {
	return c.WriteTimeout
}

func (c *Redis) GetIdleTimeout() time.Duration {
	return c.IdleTimeout
}

func (c *Redis) GetIdleCheckFrequency() time.Duration {
	return c.IdleCheckFrequency
}
