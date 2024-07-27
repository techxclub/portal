package config

import (
	"time"
)

var cfg *Config

type Config struct {
	AppName     string        `yaml:"APP_NAME" env:"APP_NAME"`
	API         HTTPAPIConfig `yaml:"API" env:",prefix=API_"`
	Admin       Admin         `yaml:"ADMIN" env:",prefix=ADMIN_"`
	Auth        Auth          `yaml:"AUTH" env:",prefix=AUTH_"`
	GoogleAuth  GoogleAuth    `yaml:"GOOGLE_AUTH" env:",prefix=GOOGLE_AUTH_"`
	Swagger     Swagger       `yaml:"SWAGGER" env:",prefix=SWAGGER_"`
	Translation Translation   `yaml:"TRANSLATION" env:",prefix=TRANSLATION_"`

	DB    DB    `yaml:"DB" env:",prefix=DB_"`
	Redis Redis `yaml:"REDIS" env:",prefix=REDIS_"`
	Log   Log   `yaml:"LOG" env:",prefix=LOG_"`
	OTP   OTP   `yaml:"OTP" env:",prefix=OTP_"`

	RateLimitEnabled bool      `yaml:"RATE_LIMIT_ENABLED" env:"RATE_LIMIT_ENABLED"`
	RateLimit        RateLimit `yaml:"RATE_LIMIT" env:",prefix=RATE_LIMIT_"`

	GoogleClient HTTPConfig `yaml:"GOOGLE_CLIENT" env:",prefix=GOOGLE_CLIENT_"`

	ServiceMail MailSMTP `yaml:"SERVICE_MAIL" env:",prefix=SERVICE_MAIL_"`
	SupportMail MailSMTP `yaml:"SUPPORT_MAIL" env:",prefix=SUPPORT_MAIL_"`

	Referral Referral `yaml:"REFERRAL" env:",prefix=REFERRAL_"`

	ResumeDirectory         string `yaml:"RESUME_DIRECTORY" env:"RESUME_DIRECTORY"`
	CompanyListLimit        int    `yaml:"COMPANY_LIST_LIMIT" env:"COMPANY_LIST_LIMIT"`
	PopularCompanyListLimit int    `yaml:"POPULAR_COMPANY_LIST_LIMIT" env:"POPULAR_COMPANY_LIST_LIMIT"`
}

type Auth struct {
	Enabled                bool          `yaml:"ENABLED" env:"ENABLED"`
	CipherKey              string        `yaml:"CIPHER_KEY" env:"CIPHER_KEY"`
	AuthIssuer             string        `yaml:"AUTH_ISSUER" env:"AUTH_ISSUER"`
	AuthIssuerUUID         string        `yaml:"AUTH_ISSUER_UUID" env:"AUTH_ISSUER_UUID"`
	AuthAudience           string        `yaml:"AUTH_AUDIENCE" env:"AUTH_AUDIENCE"`
	AccessTokenSecret      string        `yaml:"ACCESS_TOKEN_SECRET" env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string        `yaml:"REFRESH_TOKEN_SECRET" env:"REFRESH_TOKEN_SECRET"`
	AuthSoftExpiryDuration time.Duration `yaml:"AUTH_SOFT_EXPIRY_DURATION" env:"AUTH_SOFT_EXPIRY_DURATION"`
	AuthHardExpiryDuration time.Duration `yaml:"AUTH_HARD_EXPIRY_DURATION" env:"AUTH_HARD_EXPIRY_DURATION"`
}

type Admin struct {
	ClientID string `yaml:"CLIENT_ID" env:"CLIENT_ID"`
	PassKey  string `yaml:"PASS_KEY" env:"PASS_KEY"`
}

type GoogleAuth struct {
	Debug            bool   `yaml:"DEBUG" env:"DEBUG"`
	ClientState      string `yaml:"CLIENT_STATE" env:"CLIENT_STATE"`
	ClientID         string `yaml:"CLIENT_ID" env:"CLIENT_ID"`
	ClientSecret     string `yaml:"CLIENT_SECRET" env:"CLIENT_SECRET"`
	RedirectHost     string `yaml:"REDIRECT_HOST" env:"REDIRECT_HOST"`
	RedirectEndpoint string `yaml:"REDIRECT_ENDPOINT" env:"REDIRECT_ENDPOINT"`
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
	DefaultLanguage string `yaml:"DEFAULT_LANGUAGE" env:"DEFAULT_LANGUAGE"`
	JSONDirectory   string `yaml:"JSON_DIRECTORY" env:"JSON_DIRECTORY"`
	HTMLDirectory   string `yaml:"HTML_DIRECTORY" env:"HTML_DIRECTORY"`
}

type Log struct {
	Level  string `yaml:"LEVEL" env:"LEVEL"`
	Output string `yaml:"OUTPUT" env:"OUTPUT"` // Should be one of "console" | "stdout"
	Format string `yaml:"FORMAT" env:"FORMAT"`
}

type OTP struct {
	TTL            time.Duration `yaml:"TTL" env:"TTL"`
	MaxRetryCount  int           `yaml:"MAX_RETRY_COUNT" env:"MAX_RETRY_COUNT"`
	MockingEnabled bool          `yaml:"MOCKING_ENABLED" env:"MOCKING_ENABLED"`
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

	cfg.Admin = Admin{
		ClientID: "admin",
		PassKey:  "admin",
	}

	cfg.Auth = Auth{
		Enabled:                true,
		CipherKey:              "6c13b7338aa24366181369dbc6540f28",
		AuthIssuer:             "portal",
		AuthIssuerUUID:         "portal-uuid",
		AuthAudience:           "techx",
		AccessTokenSecret:      "af77aa93-42a0-4dae-add9-ade13453a770",
		RefreshTokenSecret:     "1b2a01e0-b6cc-4258-8965-d41b4bb2544d",
		AuthSoftExpiryDuration: 7 * 24 * time.Hour,
		AuthHardExpiryDuration: 30 * 24 * time.Hour,
	}

	cfg.GoogleAuth = GoogleAuth{
		Debug:            true,
		ClientState:      "google",
		ClientID:         "google-client-id",
		ClientSecret:     "google-client-secret",
		RedirectHost:     "https://localhost:5173",
		RedirectEndpoint: "/oauth2/callback",
	}

	cfg.Swagger = Swagger{
		Enabled: true,
		Path:    "./swagger",
	}

	cfg.Translation = Translation{
		DefaultLanguage: "en",
		JSONDirectory:   "./resources/json",
		HTMLDirectory:   "./resources/html",
	}

	cfg.DB = defaultDBConfig()
	cfg.Redis = defaultRedisConfig()

	cfg.RateLimitEnabled = true
	cfg.RateLimit = defaultRateLimit()

	cfg.Log = Log{
		Level:  "info",
		Output: "console",
		Format: "json",
	}

	cfg.OTP = OTP{
		TTL:            10 * time.Minute,
		MaxRetryCount:  3,
		MockingEnabled: false,
	}

	cfg.GoogleClient = HTTPConfig{
		Host:                   "www.googleapis.com",
		HTTPTimeout:            1000,
		HystrixTimeout:         1000,
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 100,
		SleepWindow:            100,
		ErrorPercentThreshold:  10,
	}

	cfg.ServiceMail = MailSMTP{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: "username",
		SMTPPassword: "password",
		Domain:       "domain.com",
		SenderEmail:  "referral@domain.com",
	}

	cfg.SupportMail = MailSMTP{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: "username",
		SMTPPassword: "password",
		Domain:       "domain.com",
		SenderEmail:  "support@domain.com",
	}

	cfg.Referral = Referral{
		RequesterReferralLimit:    20,
		ProviderReferralLimit:     10,
		ReferralMaxLookupDuration: 7 * 24 * time.Hour,
	}

	cfg.ResumeDirectory = "./user_resumes"
	cfg.CompanyListLimit = 100
	cfg.PopularCompanyListLimit = 5
}
