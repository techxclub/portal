package config

import "time"

var cfg *Config

type Config struct {
	AppName string        `yaml:"APP_NAME" env:"APP_NAME"`
	API     HTTPAPIConfig `yaml:"API" env:",prefix=API_"`
	Swagger Swagger       `yaml:"SWAGGER" env:",prefix=SWAGGER_"`
	DB      DB            `yaml:"DB" env:",prefix=DB_"`
	Log     Log           `yaml:"LOG" env:",prefix=LOG_"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}

type HTTPAPIConfig struct {
	ListenAddr string `yaml:"LISTEN_ADDR" env:"HTTP_LISTEN_ADDR"`
	DebugMode  bool   `yaml:"DEBUG_MODE" env:"DEBUG_MODE"`
}

type Log struct {
	Level  string `yaml:"LEVEL" env:"LEVEL"`
	Output string `yaml:"OUTPUT" env:"OUTPUT"` // Should be one of "console" | "stdout"
	Format string `yaml:"FORMAT" env:"FORMAT"`
}

func NewConfig(path string) (*Config, error) {
	cfg = &Config{}
	cfg.SetDefaults()

	err := LoadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) SetDefaults() {
	cfg.AppName = "portal"

	cfg.API = HTTPAPIConfig{
		ListenAddr: ":3000",
		DebugMode:  false,
	}

	cfg.Swagger = Swagger{
		Enabled: true,
		Path:    "./swagger",
	}

	cfg.DB = DB{
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

	cfg.Log = Log{
		Level:  "info",
		Output: "console",
		Format: "json",
	}
}
