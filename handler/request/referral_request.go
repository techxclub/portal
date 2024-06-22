package request

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"
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
}

func NewReferralRequest(r *http.Request) (*ReferralRequest, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return nil, errors.New("Error parsing form data")
	}
	file, _, err := r.FormFile(constants.ParamResumeFile)
	if err != nil {
		return nil, errors.New("Error retrieving the file")
	}

	resumeFilePath, err := saveResume(file)
	if err != nil {
		return nil, errors.ErrSavingResume
	}

	return &ReferralRequest{
		RequesterUserID: r.FormValue(constants.ParamRequesterID),
		ProviderUserID:  r.FormValue(constants.ParamProviderID),
		CompanyID:       utils.ParseInt64WithDefault(r.FormValue(constants.ParamCompanyID), 0),
		JobLink:         r.FormValue(constants.ParamJobLink),
		Message:         r.FormValue(constants.ParamMessage),
		ResumeFilePath:  resumeFilePath,
	}, nil
}

func (r ReferralRequest) Validate() error {
	if r.RequesterUserID == "" {
		return errors.ErrRequesterFieldIsEmpty
	}

	if r.ProviderUserID == "" {
		return errors.ErrProviderFieldIsEmpty
	}

	_, err := url.ParseRequestURI(r.JobLink)
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
		ResumeFilePath:  r.ResumeFilePath,
		Status:          constants.ReferralStatusPending,
	}
}

func saveResume(file multipart.File) (string, error) {
	tempFile, err := os.CreateTemp(os.TempDir(), "resume_*.pdf")
	if err != nil {
		log.Error().Err(err).Msg("Cannot create temporary file")
		return "", err
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write to temporary file")
		return "", err
	}

	err = tempFile.Close()
	if err != nil {
		log.Error().Err(err).Msg("Failed to close temporary file")
		return "", err
	}

	return tempFile.Name(), nil
}
