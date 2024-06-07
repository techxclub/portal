package twilio

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type Client interface {
	SendOTP(ctx context.Context, to, channel string) error
}

type client struct {
	verifyServiceSID string
	twilioClient     *twilio.RestClient
}

func NewTwilioClient(twilioConfig config.Twilio) Client {
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioConfig.AccountSID,
		Password: twilioConfig.AuthToken,
	})

	return &client{
		verifyServiceSID: twilioConfig.VerifyServiceSID,
		twilioClient:     twilioClient,
	}
}

func (c client) SendOTP(_ context.Context, to, channel string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel(channel)

	_, err := c.twilioClient.VerifyV2.CreateVerification(c.verifyServiceSID, params)
	return err
}
