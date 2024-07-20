package google

import (
	"context"

	"github.com/techx/portal/client/http"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

const (
	googleClientCmd = "googleClient"

	paramAccessToken    = "access_token"
	userInfoAPIEndpoint = "/oauth2/v2/userinfo"
)

type Client interface {
	FetchUserInfo(ctx context.Context, oauthDetails domain.GoogleOAuthDetails) (*UserInfo, error)
}

type googleClient struct {
	host string
	doer http.Doer
}

type UserInfo struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	Gender        string `json:"gender"`
	GivenName     string `json:"given_name"`
	Hd            string `json:"hd"`
	ID            string `json:"id"`
	Link          string `json:"link"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func NewGoogleClient(googleClientConfig config.HTTPConfig) Client {
	return &googleClient{
		host: googleClientConfig.Host,
		doer: http.DefaultDoer(googleClientCmd, googleClientConfig),
	}
}

func (gc *googleClient) FetchUserInfo(ctx context.Context, oauthDetails domain.GoogleOAuthDetails) (*UserInfo, error) {
	var userInfo UserInfo
	err := http.NewRequest(ctx, googleClientCmd).
		SetScheme(http.SchemeHTTPS).
		SetMethod(http.MethodGet).
		SetHost(gc.host).
		SetPath(userInfoAPIEndpoint).
		SetQueryParam(paramAccessToken, oauthDetails.AccessToken).
		Send(gc.doer, &userInfo, nil)

	return &userInfo, err
}
