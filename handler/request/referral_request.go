package request

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/utils"
)

type ReferralRequest struct {
	RequesterUserID string
	ProviderUserID  string
	CompanyID       int64
	JobLink         string
	Message         string
	ResumeFilePath  string
	ResumeFile      multipart.File
}

func NewReferralRequest(r *http.Request) (*ReferralRequest, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return nil, errors.New("Error parsing form data")
	}
	resumeFile, _, err := r.FormFile(constants.ParamResumeFile)
	if err != nil {
		return nil, errors.New("Error retrieving the resume_file")
	}

	return &ReferralRequest{
		RequesterUserID: r.FormValue(constants.ParamRequesterID),
		ProviderUserID:  r.FormValue(constants.ParamProviderID),
		CompanyID:       utils.ParseInt64WithDefault(r.FormValue(constants.ParamCompanyID), 0),
		JobLink:         r.FormValue(constants.ParamJobLink),
		Message:         r.FormValue(constants.ParamMessage),
		ResumeFile:      resumeFile,
	}, nil
}

func (r ReferralRequest) Validate() error {
	if r.RequesterUserID == "" {
		return errors.ErrRequesterFieldIsEmpty
	}

	if r.ProviderUserID == "" {
		return errors.ErrProviderFieldIsEmpty
	}

	jobURL := addSchemeIfMissing(r.JobLink)
	_, err := url.ParseRequestURI(jobURL)
	if err != nil {
		return errors.ErrInvalidJobLink
	}
	return nil
}

func (r ReferralRequest) ToReferral() domain.ReferralParams {
	return domain.ReferralParams{
		RequesterUserID: r.RequesterUserID,
		ProviderUserID:  r.ProviderUserID,
		JobLink:         r.JobLink,
		CompanyID:       r.CompanyID,
		Message:         r.Message,
		ResumeFile:      r.ResumeFile,
		Status:          constants.ReferralStatusPending,
	}
}

func addSchemeIfMissing(url string) string {
	if !strings.Contains(url, "://") {
		return "https://" + url
	}

	return url
}
