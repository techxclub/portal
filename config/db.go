package config

import (
	"fmt"
	"math/rand"
	"time"
)

type DB struct {
	Name                  string        `yaml:"NAME" env:"NAME"`
	Host                  string        `yaml:"HOST" env:"HOST"`
	Port                  int           `yaml:"PORT" env:"PORT"`
	User                  string        `yaml:"USER" env:"USER"`
	Password              string        `yaml:"PASSWORD" env:"PASSWORD"`
	SSLMode               string        `yaml:"SSL_MODE" env:"SSL_MODE"`
	MaxPoolSize           int           `yaml:"POOL" env:"POOL"`
	MaxIdleConnections    int           `yaml:"MAX_IDLE_CONNECTIONS" env:"MAX_IDLE_CONNECTIONS"`
	ConnMaxIdleTime       time.Duration `yaml:"CONNECTION_MAX_IDLE_TIME" env:"CONNECTION_MAX_IDLE_TIME"`
	ConnMaxLifeTime       time.Duration `yaml:"CONNECTION_MAX_LIFE_TIME" env:"CONNECTION_MAX_LIFE_TIME"`
	ConnMaxLifeTimeJitter time.Duration `yaml:"CONNECTION_MAX_LIFE_TIME_JITTER" env:"CONNECTION_MAX_LIFE_TIME_JITTER"`
}

func defaultDBConfig() DB {
	return DB{
		Name:                  "portal_local",
		Host:                  "localhost",
		Port:                  5432,
		User:                  "postgres",
		Password:              "",
		SSLMode:               "disable",
		MaxPoolSize:           10,
		MaxIdleConnections:    5,
		ConnMaxIdleTime:       5 * time.Minute,
		ConnMaxLifeTime:       30 * time.Minute,
		ConnMaxLifeTimeJitter: 5 * time.Minute,
	}
}

func (c *DB) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

func (c *DB) GetMaxIdleConnections() int {
	return c.MaxIdleConnections
}

func (c *DB) GetMaxPoolSize() int {
	return c.MaxPoolSize
}

func (c *DB) GetConnectionMaxIdleTime() time.Duration {
	return c.ConnMaxIdleTime
}

func (c *DB) GetConnectionMaxLifeTime() time.Duration {
	var jitter time.Duration
	if c.ConnMaxLifeTimeJitter > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		jitter = time.Duration(r.Int63n(int64(c.ConnMaxLifeTimeJitter)))
	}

	return c.ConnMaxLifeTime + jitter
}
