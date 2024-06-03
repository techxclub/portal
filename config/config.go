package config

var cfg *Config

type Config struct {
	AppName string        `yaml:"APP_NAME" env:"APP_NAME"`
	API     HTTPAPIConfig `yaml:"API" env:",prefix=API_"`
	Swagger Swagger       `yaml:"SWAGGER" env:",prefix=SWAGGER_"`

	LogLevel  string `yaml:"LOG_LEVEL" env:"LOG_LEVEL"`
	LogOutput string `yaml:"LOG_OUTPUT" env:"LOG_OUTPUT"` // Should be one of "console" | "stdout"
	LogFormat string `yaml:"LOG_FORMAT" env:"LOG_FORMAT"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}

type HTTPAPIConfig struct {
	ListenAddr string `yaml:"LISTEN_ADDR" env:"HTTP_LISTEN_ADDR"`
	DebugMode  bool   `yaml:"DEBUG_MODE" env:"DEBUG_MODE"`
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
}
