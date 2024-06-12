package builder

import (
	"context"
	"fmt"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"gopkg.in/gomail.v2"
)

type ReferralMailParams struct {
	Requester domain.UserProfile
	Provider  domain.UserProfile
	JobLink   string
	Message   string
}

type MailBuilder interface {
	SendReferralMail(ctx context.Context, params ReferralMailParams) error
}

type mailBuilder struct {
	cfg   config.Config
	GMail *gomail.Dialer
}

func NewMailBuilder(cfg config.Config, dialer *gomail.Dialer) MailBuilder {
	return &mailBuilder{
		cfg:   cfg,
		GMail: dialer,
	}
}

func (mb *mailBuilder) SendReferralMail(_ context.Context, params ReferralMailParams) error {
	subject := "TechX: Referral Request For " + params.Requester.Name
	body := fmt.Sprintf(`
	<html>
	<body>
	<p>Hey, %s</p>
	<p>We hope you're doing well!</p>
	<p>We are reaching out from the TechX community because one of our members is interested in a job opportunity at your company.</p>
	<p>Requester details:</p>
	<ul>
	  <li>Name: %s</li>
	  <li>Company: %s</li>
	  <li>Email: %s</li>
      <li>Job Link: %s</li>
	</ul>
	<p>Message from requester: %s</p>
	<p>We hope you consider their application!</p>
	<p>Best regards,</p>
	<p>The TechX Community</p>
	</body>
	</html>
	`, params.Provider.Name, params.Requester.Name, params.Requester.Company, params.Requester.PersonalEmail, params.JobLink, params.Message)
	from := mb.cfg.GMail.From
	to := []string{params.Provider.WorkEmail, params.Requester.PersonalEmail}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return mb.GMail.DialAndSend(m)
}
