package builder

import (
	"context"

	"github.com/techx/portal/apicontext"
	googleClient "github.com/techx/portal/client/google"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	keyIDToken = "id_token"

	userOpenIDScope  = "openid"
	userEmailScope   = "https://www.googleapis.com/auth/userinfo.email"
	userProfileScope = "https://www.googleapis.com/auth/userinfo.profile"
)

type GoogleOAuthBuilder interface {
	BuildGoogleOAuthDetails(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.GoogleOAuthDetails, error)
	BuildUserProfile(ctx context.Context, googleOAuthDetails domain.GoogleOAuthDetails) (*domain.User, error)
}

type googleOAuthBuilder struct {
	clientConfig config.GoogleAuth
	googleClient googleClient.Client
}

func NewGoogleOAuthBuilder(oauthConfig config.GoogleAuth, client googleClient.Client) GoogleOAuthBuilder {
	return &googleOAuthBuilder{
		clientConfig: oauthConfig,
		googleClient: client,
	}
}

func (gb googleOAuthBuilder) BuildGoogleOAuthDetails(ctx context.Context, exchangeReq domain.GoogleOAuthExchangeRequest) (*domain.GoogleOAuthDetails, error) {
	oauthCode := exchangeReq.Code
	if exchangeReq.OAuthCode != "" {
		oauthCode = exchangeReq.OAuthCode
	}
	token, err := gb.getOAuthConfig(ctx).Exchange(ctx, oauthCode)
	if err != nil {
		return nil, err
	}

	return &domain.GoogleOAuthDetails{
		TokenType:    token.TokenType,
		AccessToken:  token.AccessToken,
		IDToken:      token.Extra(keyIDToken).(string),
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}, nil
}

func (gb googleOAuthBuilder) BuildUserProfile(ctx context.Context, googleOAuthDetails domain.GoogleOAuthDetails) (*domain.User, error) {
	googleUserInfo, err := gb.googleClient.FetchUserInfo(ctx, googleOAuthDetails)
	if err != nil {
		return nil, err
	}

	if googleUserInfo.Email == "" || !googleUserInfo.VerifiedEmail {
		return nil, errors.ErrUnverifiedEmail
	}

	userProfile := &domain.User{
		Status: constants.StatusIncompleteProfile,
		PersonalInformation: domain.PersonalInformation{
			Name:            googleUserInfo.FullName(),
			RegisteredEmail: googleUserInfo.Email,
			ProfilePicture:  googleUserInfo.Picture,
		},
	}
	googleOAuthDetails.Email = googleUserInfo.Email
	userProfile.SetGoogleOAuthDetails(googleOAuthDetails)
	return userProfile, nil
}

func (gb googleOAuthBuilder) getOAuthConfig(ctx context.Context) *oauth2.Config {
	redirectHost := gb.clientConfig.RedirectHost
	if origin := apicontext.RequestContextFromContext(ctx).GetOrigin(); origin != "" {
		redirectHost = origin
	}

	redirectURL := redirectHost + gb.clientConfig.RedirectEndpoint
	return &oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     gb.clientConfig.ClientID,
		ClientSecret: gb.clientConfig.ClientSecret,
		Scopes: []string{
			userOpenIDScope,
			userEmailScope,
			userProfileScope,
		},
		Endpoint: google.Endpoint,
	}
}
