package config

import (
	"time"

	"github.com/techx/portal/constants"
)

var cfg *Config

type Config struct {
	AppName     string        `yaml:"APP_NAME" env:"APP_NAME"`
	API         HTTPAPIConfig `yaml:"API" env:",prefix=API_"`
	Auth        *Auth         `yaml:"AUTH" env:",prefix=AUTH_"`
	AdminAuth   *AdminAuth    `yaml:"ADMIN_AUTH" env:",prefix=ADMIN_AUTH_"`
	Swagger     Swagger       `yaml:"SWAGGER" env:",prefix=SWAGGER_"`
	Translation Translation   `yaml:"TRANSLATION" env:",prefix=TRANSLATION_"`

	DB       DB       `yaml:"DB" env:",prefix=DB_"`
	Redis    Redis    `yaml:"REDIS" env:",prefix=REDIS_"`
	Log      Log      `yaml:"LOG" env:",prefix=LOG_"`
	OTP      OTP      `yaml:"OTP" env:",prefix=OTP_"`
	Gmail    Gmail    `yaml:"GMAIL" env:",prefix=GMAIL_"`
	Referral Referral `yaml:"REFERRAL" env:",prefix=REFERRAL_"`

	CompanyListLimit        int `yaml:"COMPANY_LIST_LIMIT" env:"COMPANY_LIST_LIMIT"`
	PopularCompanyListLimit int `yaml:"POPULAR_COMPANY_LIST_LIMIT" env:"POPULAR_COMPANY_LIST_LIMIT"`
}

type Auth struct {
	DebugMode              bool          `yaml:"DEBUG_MODE" env:"DEBUG_MODE"`
	CipherKey              string        `yaml:"CIPHER_KEY" env:"CIPHER_KEY"`
	AuthIssuer             string        `yaml:"AUTH_ISSUER" env:"AUTH_ISSUER"`
	AuthIssuerUUID         string        `yaml:"AUTH_ISSUER_UUID" env:"AUTH_ISSUER_UUID"`
	AuthAudience           string        `yaml:"AUTH_AUDIENCE" env:"AUTH_AUDIENCE"`
	AccessTokenSecret      string        `yaml:"ACCESS_TOKEN_SECRET" env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string        `yaml:"REFRESH_TOKEN_SECRET" env:"REFRESH_TOKEN_SECRET"`
	AuthSoftExpiryDuration time.Duration `yaml:"AUTH_SOFT_EXPIRY_DURATION" env:"AUTH_SOFT_EXPIRY_DURATION"`
	AuthHardExpiryDuration time.Duration `yaml:"AUTH_HARD_EXPIRY_DURATION" env:"AUTH_HARD_EXPIRY_DURATION"`
}

type AdminAuth struct {
	ClientID string `yaml:"CLIENT_ID" env:"CLIENT_ID"`
	PassKey  string `yaml:"PASS_KEY" env:"PASS_KEY"`
}

type HTTPAPIConfig struct {
	ListenAddr string `yaml:"LISTEN_ADDR" env:"HTTP_LISTEN_ADDR"`
	DebugMode  bool   `yaml:"DEBUG_MODE" env:"DEBUG_MODE"`
}

type Swagger struct {
	Enabled bool   `yaml:"ENABLED" env:"ENABLED"`
	Path    string `yaml:"PATH" env:"PATH"`
}

type Translation struct {
	FilePath        string `yaml:"FILE_PATH" env:"FILE_PATH"`
	DefaultLanguage string `yaml:"DEFAULT_LANGUAGE" env:"DEFAULT_LANGUAGE"`
}

type Log struct {
	Level  string `yaml:"LEVEL" env:"LEVEL"`
	Output string `yaml:"OUTPUT" env:"OUTPUT"` // Should be one of "console" | "stdout"
	Format string `yaml:"FORMAT" env:"FORMAT"`
}

type OTP struct {
	TTL                     time.Duration `yaml:"TTL" env:"TTL"`
	MaxRetryCount           int           `yaml:"MAX_RETRY_COUNT" env:"MAX_RETRY_COUNT"`
	MockingEnabled          bool          `yaml:"MOCKING_ENABLED" env:"MOCKING_ENABLED"`
	EmailThirdPartyProvider string        `yaml:"EMAIL_THIRD_PARTY_PROVIDER" env:"EMAIL_THIRD_PARTY_PROVIDER"`
	SMSThirdPartyProvider   string        `yaml:"SMS_THIRD_PARTY_PROVIDER" env:"SMS_THIRD_PARTY_PROVIDER"`
}

type Gmail struct {
	SMTPServer   string `yaml:"SMTP_SERVER" env:"SMTP_SERVER"`
	SMTPPort     int    `yaml:"SMTP_PORT" env:"SMTP_PORT"`
	SMTPUsername string `yaml:"SMTP_USERNAME" env:"SMTP_USERNAME"`
	SMTPPassword string `yaml:"SMTP_PASSWORD" env:"SMTP_PASSWORD"`
	From         string `yaml:"FROM" env:"FROM"`
}

type Referral struct {
	RequesterReferralLimit    int           `yaml:"REQUESTER_REFERRAL_LIMIT" env:"REQUESTER_REFERRAL_LIMIT"`
	ProviderReferralLimit     int           `yaml:"PROVIDER_REFERRAL_LIMIT" env:"PROVIDER_REFERRAL_LIMIT"`
	ReferralMaxLookupDuration time.Duration `yaml:"REFERRAL_MAX_LOOKUP_DURATION" env:"REFERRAL_MAX_LOOKUP_DURATION"`
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

	cfg.Auth = &Auth{
		DebugMode:              false,
		CipherKey:              "6c13b7338aa24366181369dbc6540f28",
		AuthIssuer:             "portal",
		AuthIssuerUUID:         "portal-uuid",
		AuthAudience:           "techx",
		AccessTokenSecret:      "af77aa93-42a0-4dae-add9-ade13453a770",
		RefreshTokenSecret:     "1b2a01e0-b6cc-4258-8965-d41b4bb2544d",
		AuthSoftExpiryDuration: 7 * 24 * time.Hour,
		AuthHardExpiryDuration: 30 * 24 * time.Hour,
	}

	cfg.AdminAuth = &AdminAuth{
		ClientID: "admin",
		PassKey:  "admin",
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

	cfg.Redis = Redis{
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

	cfg.Log = Log{
		Level:  "info",
		Output: "console",
		Format: "json",
	}

	cfg.OTP = OTP{
		TTL:                     10 * time.Minute,
		MaxRetryCount:           3,
		MockingEnabled:          false,
		EmailThirdPartyProvider: constants.ThirdPartyGomail,
		SMSThirdPartyProvider:   constants.ThirdPartyMsg91,
	}

	cfg.Gmail = Gmail{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: "username",
		SMTPPassword: "password",
		From:         "user.name@gmail.com",
	}

	cfg.Referral = Referral{
		RequesterReferralLimit:    20,
		ProviderReferralLimit:     10,
		ReferralMaxLookupDuration: 7 * 24 * time.Hour,
	}

	cfg.CompanyListLimit = 100
	cfg.PopularCompanyListLimit = 5
}
