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
	RequesterUserUUID string
	ProviderUserUUID  string
	CompanyID         int64
	JobLink           string
	Message           string
	ResumeFilePath    string
	ResumeFile        multipart.File
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
		RequesterUserUUID: r.FormValue(constants.ParamRequesterID),
		ProviderUserUUID:  r.FormValue(constants.ParamProviderID),
		CompanyID:         utils.ParseInt64WithDefault(r.FormValue(constants.ParamCompanyID), 0),
		JobLink:           r.FormValue(constants.ParamJobLink),
		Message:           r.FormValue(constants.ParamMessage),
		ResumeFile:        resumeFile,
	}, nil
}

func (r ReferralRequest) Validate() error {
	if r.RequesterUserUUID == "" {
		return errors.ErrRequesterFieldIsEmpty
	}

	if r.ProviderUserUUID == "" {
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
		Referral: domain.Referral{
			Status:            constants.ReferralStatusPending,
			RequesterUserUUID: r.RequesterUserUUID,
			ProviderUserUUID:  r.ProviderUserUUID,
			JobLink:           r.JobLink,
			CompanyID:         r.CompanyID,
		},
		Message:    r.Message,
		ResumeFile: r.ResumeFile,
	}
}

func addSchemeIfMissing(url string) string {
	if !strings.Contains(url, "://") {
		return "https://" + url
	}

	return url
}
