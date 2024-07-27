package builder

import (
	"context"

	"github.com/techx/portal/client/email"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/i18n"
	"gopkg.in/gomail.v2"
)

const (
	jobNameRequestReferral  = "referral_mail"
	jobNameSendOTP          = "send_otp"
	senderNameTechXReferral = "TechX"
	senderNameTechXSupport  = "TechX"
	senderNameTechXService  = "TechX"

	i18nKeyEmailOTPMailSubject = "otp_mail_subject"
	i18nFileEmailOTPMailBody   = "otp_mail_body"

	i18nKeyApprovalMailSubject = "user_approval_mail_subject"
	i18nFileApprovalMailBody   = "user_approval_mail_body"

	i18nKeyReferralMailSubject                = "referral_mail_subject"
	i18nFileReferralMailBodyHTML              = "referral_mail_body"
	i18nFileReferralMailCustomMessageBodyHTML = "referral_mail_with_custom_message_body"
)

type OTPMailParams struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

type ReferralMailParams struct {
	Requester         domain.User `json:"requester"`
	Provider          domain.User `json:"provider"`
	NoticePeriod      string      `json:"notice_period"`
	PreferredLocation string      `json:"preferred_location"`
	JobLink           string      `json:"job_link"`
	Message           string      `json:"message"`
	ResumeFilePath    string      `json:"resume_file_path"`
	AttachmentName    string      `json:"attachment_name"`
}

type ApprovalMailParams struct {
	User domain.User `json:"user"`
}

type MailBuilder interface {
	SendOTPMail(ctx context.Context, async bool, refID string, otpMailParams OTPMailParams) error
	SendUserApprovalMail(ctx context.Context, async bool, refID string, approvalMailParams ApprovalMailParams) error
	SendReferralMail(ctx context.Context, async bool, refID string, referralMailParams ReferralMailParams) error
}

type mailBuilder struct {
	serviceMailClient email.Client
	supportMailClient email.Client
}

func NewMailBuilder(serviceMailClient, supportMailClient email.Client) MailBuilder {
	return &mailBuilder{
		serviceMailClient: serviceMailClient,
		supportMailClient: supportMailClient,
	}
}

func (mb *mailBuilder) SendOTPMail(ctx context.Context, async bool, refID string, otpMailParams OTPMailParams) error {
	i18nValues := map[string]interface{}{
		"OTP": otpMailParams.Code,
	}

	subject := i18n.Translate(ctx, i18nKeyEmailOTPMailSubject)
	bodyHTML := i18n.HTML(ctx, i18nFileEmailOTPMailBody, i18nValues)

	message := gomail.NewMessage()
	message.SetHeader(constants.GomailHeaderTo, otpMailParams.Email)
	message.SetHeader(constants.GomailHeaderSubject, subject)
	message.SetHeader(constants.GomailHeaderMessageID, refID)
	message.SetHeader(constants.GomailHeaderInReplyTo, refID)
	message.SetHeader(constants.GomailHeaderReferences, refID)
	message.SetBody(constants.GomailContentTypeHTML, bodyHTML)

	return sendMail(ctx, mb.supportMailClient, async, jobNameSendOTP, senderNameTechXSupport, message)
}

func (mb *mailBuilder) SendUserApprovalMail(ctx context.Context, async bool, refID string, approvalMailParams ApprovalMailParams) error {
	i18nValues := map[string]interface{}{
		"UserName": approvalMailParams.User.Name,
	}

	subject := i18n.Translate(ctx, i18nKeyApprovalMailSubject, i18nValues)
	bodyHTML := i18n.HTML(ctx, i18nFileApprovalMailBody, i18nValues)
	to := []string{approvalMailParams.User.RegisteredEmail}

	message := gomail.NewMessage()
	message.SetHeader(constants.GomailHeaderTo, to...)
	message.SetHeader(constants.GomailHeaderSubject, subject)
	message.SetHeader(constants.GomailHeaderMessageID, refID)
	message.SetHeader(constants.GomailHeaderInReplyTo, refID)
	message.SetHeader(constants.GomailHeaderReferences, refID)
	message.SetBody(constants.GomailContentTypeHTML, bodyHTML)

	return sendMail(ctx, mb.serviceMailClient, async, jobNameRequestReferral, senderNameTechXService, message)
}

func (mb *mailBuilder) SendReferralMail(ctx context.Context, async bool, refID string, referralMailParams ReferralMailParams) error {
	i18nValues := map[string]interface{}{
		"ProviderName":         referralMailParams.Provider.Name,
		"RequesterName":        referralMailParams.Requester.Name,
		"RequesterCompany":     referralMailParams.Requester.CompanyName,
		"RequesterDesignation": referralMailParams.Requester.Designation,
		"RequesterYOE":         referralMailParams.Requester.YearsOfExperience,
		"RequesterEmail":       referralMailParams.Requester.RegisteredEmail,
		"NoticePeriod":         referralMailParams.NoticePeriod,
		"PreferredLocation":    referralMailParams.PreferredLocation,
		"JobLink":              referralMailParams.JobLink,
		"RequesterMessage":     referralMailParams.Message,
	}

	subject := i18n.Translate(ctx, i18nKeyReferralMailSubject, i18nValues)
	bodyHTMLFile := i18nFileReferralMailBodyHTML
	if referralMailParams.Message != "" {
		bodyHTMLFile = i18nFileReferralMailCustomMessageBodyHTML
	}
	bodyHTMLString := i18n.HTML(ctx, bodyHTMLFile, i18nValues)
	to := []string{referralMailParams.Provider.WorkEmail, referralMailParams.Requester.RegisteredEmail}

	message := gomail.NewMessage()
	message.SetHeader(constants.GomailHeaderTo, to...)
	message.SetHeader(constants.GomailHeaderSubject, subject)
	message.SetHeader(constants.GomailHeaderMessageID, refID)
	message.SetHeader(constants.GomailHeaderInReplyTo, refID)
	message.SetHeader(constants.GomailHeaderReferences, refID)
	message.SetBody(constants.GomailContentTypeHTML, bodyHTMLString)
	message.Attach(referralMailParams.ResumeFilePath, gomail.Rename(referralMailParams.AttachmentName))

	return sendMail(ctx, mb.serviceMailClient, async, jobNameRequestReferral, senderNameTechXReferral, message)
}

func sendMail(ctx context.Context, mailClient email.Client, async bool, jobName, senderName string, message *gomail.Message) error {
	if async {
		mailClient.SendEmailAsync(ctx, jobName, senderName, message)
		return nil
	}

	return mailClient.SendEmail(ctx, senderName, message)
}
