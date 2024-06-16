package builder

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/i18n"
	"gopkg.in/gomail.v2"
)

const (
	i18nKeyReferralMailSubject           = "referral_mail_subject"
	i18nKeyReferralMailBodyBase          = "referral_mail_body_base"
	i18nKeyReferralMailBodyCustomMessage = "referral_mail_body_requester_customer_message"
	i18nKeyReferralMailBodyFooterNotes   = "referral_mail_body_footer_notes"
)

type ReferralMailParams struct {
	Requester      domain.UserProfile
	Provider       domain.UserProfile
	JobLink        string
	Message        string
	ResumeFilePath string
}

type MailBuilder interface {
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

func (mb *mailBuilder) SendReferralMail(ctx context.Context, params ReferralMailParams) error {
	i18nValues := map[string]interface{}{
		"ProviderName":     params.Provider.Name,
		"RequesterName":    params.Requester.Name,
		"RequesterCompany": params.Requester.Company,
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

	from := mb.cfg.GMail.From
	to := []string{params.Provider.WorkEmail, params.Requester.PersonalEmail}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", textHTML)

	defer os.Remove(params.ResumeFilePath)

	m.Attach(params.ResumeFilePath, gomail.Rename(getResumeFileName(params.Requester.Name)))

	return mb.GMail.DialAndSend(m)
}

func getResumeFileName(name string) string {
	temp := strings.Split(name, " ")
	sanitizedName := strings.Join(temp, "_")
	return fmt.Sprintf("Resume_%s.pdf", sanitizedName)
}
