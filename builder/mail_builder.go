package builder

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/i18n"
	"gopkg.in/gomail.v2"
)

const (
	failedJobsDir = "./failed_jobs/"

	i18nKeyReferralMailSubject           = "referral_mail_subject"
	i18nKeyReferralMailBodyBase          = "referral_mail_body_base"
	i18nKeyReferralMailBodyCustomMessage = "referral_mail_body_requester_customer_message"
	i18nKeyReferralMailBodyFooterNotes   = "referral_mail_body_footer_notes"
)

type ReferralMailParams struct {
	Requester      domain.UserProfile `json:"requester"`
	Provider       domain.UserProfile `json:"provider"`
	JobLink        string             `json:"job_link"`
	Message        string             `json:"message"`
	ResumeFilePath string             `json:"resume_file_path"`
}

type MailBuilder interface {
	SendReferralMailAsync(ctx context.Context, params ReferralMailParams)
	SendReferralMail(ctx context.Context, params ReferralMailParams) error
}

type mailBuilder struct {
	cfg   *config.Config
	GMail *gomail.Dialer
}

func NewMailBuilder(cfg *config.Config, dialer *gomail.Dialer) MailBuilder {
	return &mailBuilder{
		cfg:   cfg,
		GMail: dialer,
	}
}

func (mb *mailBuilder) SendReferralMailAsync(ctx context.Context, params ReferralMailParams) {
	go func() {
		maxRetries := 5
		var err error
		for i := 0; i < maxRetries; i++ {
			log.Info().Msgf("Start processing referral mail with requestor_user_id: %s, provider_user_id: %s and "+
				"filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)
			err = mb.SendReferralMail(ctx, params)
			log.Info().Msgf("Finish processing referral mail with requestor_user_id: %s, provider_user_id: %s and "+
				"filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)

			if err != nil {
				log.Printf("Failed to send email: %v", err)
				time.Sleep(time.Second * 2) // wait for 2 seconds before retrying
				continue
			}
			break
		}

		if err == nil {
			return
		}
		log.Error().Err(err).Msgf("Failed to send referral mail with requestor_user_id: %s, "+
			"provider_user_id: %s and filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)

		err = storeFailedJob(params)
		if err != nil {
			log.Printf("Failed to store failed job: %v", err)
		}
	}()
}

func (mb *mailBuilder) SendReferralMail(ctx context.Context, params ReferralMailParams) error {
	i18nValues := map[string]interface{}{
		"ProviderName":     params.Provider.Name,
		"RequesterName":    params.Requester.Name,
		"RequesterCompany": params.Requester.CompanyName,
		"RequesterEmail":   params.Requester.PersonalEmail,
		"JobLink":          params.JobLink,
		"RequesterMessage": params.Message,
	}

	subject := i18n.Translate(ctx, i18nKeyReferralMailSubject, i18nValues)
	bodyHTML := i18n.Translate(ctx, i18nKeyReferralMailBodyBase, i18nValues)
	if params.Message != "" {
		bodyHTML += i18n.Translate(ctx, i18nKeyReferralMailBodyCustomMessage, i18nValues)
	}
	bodyHTML += i18n.Translate(ctx, i18nKeyReferralMailBodyFooterNotes, i18nValues)

	textHTML := fmt.Sprintf(`<html><body>%s</body></html>`, bodyHTML)

	from := mb.cfg.Gmail.From
	to := []string{params.Provider.WorkEmail, params.Requester.PersonalEmail}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", textHTML)
	m.Attach(params.ResumeFilePath, gomail.Rename(getResumeFileName(params.Requester.Name)))

	err := mb.GMail.DialAndSend(m)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send referral mail with requestor_user_id: %s, "+
			"provider_user_id: %s and filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)
		return err
	}

	if err := os.Remove(params.ResumeFilePath); err != nil {
		log.Error().Err(err).Msgf("Failed to delete resume file with requestor_user_id: %s, "+
			"provider_user_id: %s and filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)
	}

	return nil
}

func getResumeFileName(name string) string {
	temp := strings.Split(name, " ")
	sanitizedName := strings.Join(temp, "_")
	return fmt.Sprintf("Resume_%s.pdf", sanitizedName)
}

func storeFailedJob(params ReferralMailParams) error {
	// Check if the directory exists
	if _, err := os.Stat(failedJobsDir); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.MkdirAll(failedJobsDir, 0o755)
		if err != nil {
			log.Printf("Failed to create directory: %v", err)
			return err
		}
	}

	failedJobJSON, _ := json.Marshal(params)
	return os.WriteFile(fmt.Sprintf("%s%s.json", failedJobsDir, params.Requester.UserID), failedJobJSON, 0o600)
}
