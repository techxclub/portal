package email

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/logger"
	"github.com/techx/portal/utils"
	"gopkg.in/gomail.v2"
)

const (
	failedJobsDir = "./failed_jobs"
)

type Client interface {
	SendEmail(ctx context.Context, senderName string, message *gomail.Message) error
	SendEmailAsync(ctx context.Context, jobName, senderName string, message *gomail.Message)
}

type emailClient struct {
	mailConfig config.MailSMTP
	mailDialer *gomail.Dialer
}

func NewEmailClient(mailConfig config.MailSMTP) Client {
	return &emailClient{
		mailConfig: mailConfig,
		mailDialer: newMailDialer(mailConfig),
	}
}

func (ec *emailClient) SendEmail(_ context.Context, senderName string, message *gomail.Message) error {
	message.SetHeader(constants.GomailHeaderFrom, ec.mailConfig.GetSender(senderName))
	err := ec.mailDialer.DialAndSend(message)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send email with message: %v", message)
	}

	return err
}

func (ec *emailClient) SendEmailAsync(ctx context.Context, jobName, senderName string, message *gomail.Message) {
	traceID := apicontext.RequestContextFromContext(ctx).GetTraceID()

	go func() {
		maxRetries := ec.mailConfig.Retries
		var err error
		for count := 0; count < maxRetries; count++ {
			err = ec.SendEmail(ctx, senderName, message)
			if err == nil {
				return
			}

			logFields := map[string]interface{}{
				"RetryCount":          count,
				logger.JobNameField:   jobName,
				logger.RequestTraceID: traceID,
				logger.ErrorField:     err.Error(),
				logger.MessageField:   message,
			}
			log.Info().Fields(logFields).Msg("Failed to send email, retrying")
		}

		err = ec.storeFailedMessage(jobName, traceID, message)
		if err != nil {
			log.Err(err).Msg("Failed to store failed job")
		}
	}()
}

func (ec *emailClient) storeFailedMessage(jobName, traceID string, message *gomail.Message) error {
	if err := utils.CreateDirectoryIfNotExist(failedJobsDir); err != nil {
		return err
	}

	// Custom serialization of gomail.Message
	failedJobJSON, err := json.Marshal(struct {
		From    string   `json:"from"`
		To      []string `json:"to"`
		Subject string   `json:"subject"`
	}{
		From:    message.GetHeader(constants.GomailHeaderFrom)[0],
		To:      message.GetHeader(constants.GomailHeaderTo),
		Subject: message.GetHeader(constants.GomailHeaderSubject)[0],
	})
	if err != nil {
		log.Err(err).Msg("Error serializing failed job payload")
	}

	return os.WriteFile(fmt.Sprintf("%s/%s-%s.json", failedJobsDir, jobName, traceID), failedJobJSON, 0o600)
}

func newMailDialer(gmailCfg config.MailSMTP) *gomail.Dialer {
	return gomail.NewDialer(
		gmailCfg.SMTPServer,
		gmailCfg.SMTPPort,
		gmailCfg.SenderEmail,
		gmailCfg.SMTPPassword,
	)
}
