package twilio

import (
	"context"

	"github.com/techx/portal/config"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

const (
	statusPending  = "pending"
	statusApproved = "approved"
)

type Client interface {
	SendOTP(ctx context.Context, req CreateVerificationRequest) (CreateVerificationResponse, error)
	VerifyOTP(ctx context.Context, req CheckVerificationRequest) (CheckVerificationResponse, error)
}

type client struct {
	twilioClient     *twilio.RestClient
	verifyServiceSID string
}

func NewTwilioClient(twilioConfig config.Twilio) Client {
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioConfig.AccountSID,
		Password: twilioConfig.AuthToken,
	})

	return &client{
		twilioClient:     twilioClient,
		verifyServiceSID: twilioConfig.VerifyServiceSID,
	}
}

func (c client) SendOTP(_ context.Context, req CreateVerificationRequest) (CreateVerificationResponse, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(req.To)
	params.SetChannel(req.Channel)

	resp, err := c.twilioClient.VerifyV2.CreateVerification(c.verifyServiceSID, params)
	if err != nil {
		return CreateVerificationResponse{}, err
	}

	if resp == nil || resp.Status == nil || *resp.Status != statusPending {
		return CreateVerificationResponse{}, errors.ErrOTPGenerateFailed
	}

	return CreateVerificationResponse{
		To:               utils.FromPtr(resp.To),
		Channel:          utils.FromPtr(resp.Channel),
		Status:           utils.FromPtr(resp.Status),
		Lookup:           resp.Lookup,
		SendCodeAttempts: *resp.SendCodeAttempts,
		DateCreated:      resp.DateCreated,
		DateUpdated:      resp.DateUpdated,
		URL:              utils.FromPtr(resp.Url),
	}, nil
}

func (c client) VerifyOTP(_ context.Context, req CheckVerificationRequest) (CheckVerificationResponse, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(req.From)
	params.SetCode(req.OTP)

	resp, err := c.twilioClient.VerifyV2.CreateVerificationCheck(c.verifyServiceSID, params)
	if err != nil {
		return CheckVerificationResponse{}, err
	}

	if resp == nil || resp.Status == nil || *resp.Status != statusApproved {
		return CheckVerificationResponse{}, errors.ErrOTPVerificationFailed
	}
	return CheckVerificationResponse{
		To:                    utils.FromPtr(resp.To),
		Channel:               utils.FromPtr(resp.Channel),
		Status:                utils.FromPtr(resp.Status),
		DateCreated:           resp.DateCreated,
		DateUpdated:           resp.DateUpdated,
		SnaAttemptsErrorCodes: resp.SnaAttemptsErrorCodes,
	}, nil
}
