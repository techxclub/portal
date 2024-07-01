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
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/i18n"
	"github.com/techx/portal/utils"
	"gopkg.in/gomail.v2"
)

const (
	failedJobsDir       = "./failed_jobs/"
	referralMailJobName = "referral_mail"

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

type ReferralMailBuilder interface {
	SendReferralMailAsync(ctx context.Context, params ReferralMailParams)
	SendReferralMail(ctx context.Context, params ReferralMailParams) error
}

type referralMailBuilder struct {
	cfg                *config.Config
	referralMailClient *gomail.Dialer
}

func NewReferralMailBuilder(cfg *config.Config, dialer *gomail.Dialer) ReferralMailBuilder {
	return &referralMailBuilder{
		cfg:                cfg,
		referralMailClient: dialer,
	}
}

func (mb *referralMailBuilder) SendReferralMailAsync(ctx context.Context, params ReferralMailParams) {
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

		err = storeFailedJob(referralMailJobName, params)
		if err != nil {
			log.Printf("Failed to store failed job: %v", err)
		}
	}()
}

func (mb *referralMailBuilder) SendReferralMail(ctx context.Context, params ReferralMailParams) error {
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
	to := []string{params.Provider.WorkEmail, params.Requester.PersonalEmail}
	mailCfg := mb.cfg.ReferralMail
	messageID := mailCfg.GetMessageID()

	m := gomail.NewMessage()
	m.SetHeader(constants.GomailHeaderFrom, mailCfg.GetFrom())
	m.SetHeader(constants.GomailHeaderTo, to...)
	m.SetHeader(constants.GomailHeaderSubject, subject)
	m.SetHeader(constants.GomailHeaderMessageID, messageID)
	m.SetHeader(constants.GomailHeaderInReplyTo, messageID)
	m.SetHeader(constants.GomailHeaderReferences, messageID)
	m.SetBody(constants.GomailContentTypeHTML, textHTML)
	m.Attach(params.ResumeFilePath, gomail.Rename(getResumeFileName(params.Requester.Name)))

	err := mb.referralMailClient.DialAndSend(m)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send referral mail with requestor_user_id: %s, "+
			"provider_user_id: %s and filepath: %s", params.Requester.UserID, params.Provider.UserID, params.ResumeFilePath)
		return err
	}

	return nil
}

func getResumeFileName(name string) string {
	temp := strings.Split(name, " ")
	sanitizedName := strings.Join(temp, "_")
	return fmt.Sprintf("Resume_%s.pdf", sanitizedName)
}

func storeFailedJob(jobName string, params ReferralMailParams) error {
	if err := utils.CreateDirectoryIfNotExist(failedJobsDir); err != nil {
		return err
	}

	failedJobJSON, _ := json.Marshal(params)
	return os.WriteFile(fmt.Sprintf("%s%s_%d_%d.json", failedJobsDir, jobName, params.Requester.UserIDNum, params.Provider.UserIDNum), failedJobJSON, 0o600)
}
