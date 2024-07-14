package builder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"gopkg.in/gomail.v2"
)

type GmailAuthBuilder interface {
	FetchAccessAndIdTokens(ctx context.Context, params domain.GmailAuthRequest) error
}

type gmailAuthBuilder struct {
	cfg                *config.Config
	referralMailClient *gomail.Dialer
}

func NewGmailAuthBuilder(cfg *config.Config, dialer *gomail.Dialer) GmailAuthBuilder {
	return &gmailAuthBuilder{
		cfg:                cfg,
		referralMailClient: dialer,
	}
}

func (mb *gmailAuthBuilder) FetchAccessAndIdTokens(ctx context.Context, params domain.GmailAuthRequest) error {
	URL := "https://oauth2.googleapis.com/token"
	resp, err := http.PostForm(URL,
		url.Values{
			"code":          {params.AuthCode},
			"client_id":     {mb.cfg.GoogleAuthClientID},
			"client_secret": {mb.cfg.GoogleAuthClientSecret},
			"redirect_uri":  {mb.cfg.RedirectURI},
			"grant_type":    {"authorization_code"},
		})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("google api: got http status %v", resp.StatusCode)
	}

	// TODO: Store the access and id tokens somewhere
	// result["access_token"], result["id_token"]

	return nil
}
