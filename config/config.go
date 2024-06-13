package config

import (
	"time"
)

var cfg *Config

type Config struct {
	AppName     string        `yaml:"APP_NAME" env:"APP_NAME"`
	API         HTTPAPIConfig `yaml:"API" env:",prefix=API_"`
	Swagger     Swagger       `yaml:"SWAGGER" env:",prefix=SWAGGER_"`
	Translation Translation   `yaml:"TRANSLATION" env:",prefix=TRANSLATION_"`

	DB       DB       `yaml:"DB" env:",prefix=DB_"`
	Log      Log      `yaml:"LOG" env:",prefix=LOG_"`
	Twilio   Twilio   `yaml:"TWILIO" env:",prefix=TWILIO_"`
	GMail    GMail    `yaml:"GMAIL" env:",prefix=GMAIL_"`
	Referral Referral `yaml:"REFERRAL" env:",prefix=REFERRAL_"`

	ThirdPartySmsProvider string `yaml:"THIRD_PARTY_SMS_PROVIDER" env:"THIRD_PARTY_SMS_PROVIDER"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}

type Translation struct {
	FilePath        string `yaml:"FILE_PATH" env:"FILE_PATH"`
	DefaultLanguage string `yaml:"DEFAULT_LANGUAGE" env:"DEFAULT_LANGUAGE"`
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

type Twilio struct {
	AccountSID       string `yaml:"ACCOUNT_SID" env:"ACCOUNT_SID"`
	AuthToken        string `yaml:"AUTH_TOKEN" env:"AUTH_TOKEN"`
	VerifyServiceSID string `yaml:"VERIFY_SERVICE_SID" env:"VERIFY_SID"`
}

type GMail struct {
	SMTPServer   string `yaml:"SMTP_SERVER" env:"SMTP_SERVER"`
	SMTPPort     int    `yaml:"SMTP_PORT" env:"SMTP_PORT"`
	SMTPUsername string `yaml:"SMTP_USERNAME" env:"SMTP_USERNAME"`
	SMTPPassword string `yaml:"SMTP_PASSWORD" env:"SMTP_PASSWORD"`
	From         string `yaml:"FROM" env:"FROM"`
}

type Referral struct {
	RequesterReferralLimit int       `yaml:"REQUESTER_REFERRAL_LIMIT" env:"REQUESTER_REFERRAL_LIMIT"`
	ProviderReferralLimit  int       `yaml:"PROVIDER_REFERRAL_LIMIT" env:"PROVIDER_REFERRAL_LIMIT"`
	ReferralMaxTime        time.Time `yaml:"REFERRAL_MAX_TIME" env:"REFERRAL_MAX_TIME"`
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

	cfg.Translation = Translation{
		FilePath:        "./i18n/definitions",
		DefaultLanguage: "en",
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

	cfg.Twilio = Twilio{
		AccountSID:       "your_account_sid",
		AuthToken:        "your_auth_token",
		VerifyServiceSID: "verify_service_sid",
	}

	cfg.ThirdPartySmsProvider = "twilio"

	cfg.GMail = GMail{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: "username",
		SMTPPassword: "password",
		From:         "user.name@gmail.com",
	}

	cfg.Referral = Referral{
		RequesterReferralLimit: 20,
		ProviderReferralLimit:  10,
		ReferralMaxTime:        time.Now().Add(-7 * 24 * time.Hour),
	}
}
