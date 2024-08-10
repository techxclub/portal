package admin

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/techx/portal/client"
	"github.com/techx/portal/client/azure"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/response"
)

func FetchCompanyLogoHandler(_ *config.Config, cr *client.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyID, err := strconv.ParseInt(r.URL.Query().Get(constants.ParamCompanyID), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req := azure.FetchImageParams{
			CompanyID: companyID,
		}
		images, err := cr.AzureStorage.FetchLogos(r.Context(), req)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(errors.ErrGeneratingAuthToken))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(images) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var resp domain.LogoList
		resp.Images = append(resp.Images, images...)

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func UploadCompanyLogoHandler(_ *config.Config, cr *client.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyID, err := strconv.ParseInt(r.FormValue(constants.ParamCompanyID), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		companyName := r.FormValue(constants.ParamCompanyName)
		if companyName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var file multipart.File
		var header *multipart.FileHeader

		file, header, err = r.FormFile(constants.ParamCompanyLogoFile)
		if err != nil {
			imageLink := r.FormValue(constants.ParamCompanyLogoLink)
			if imageLink == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			resp, err := http.Get(imageLink) //nolint
			if err != nil || resp.StatusCode != http.StatusOK {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer resp.Body.Close()

			tempFile, err := os.CreateTemp(os.TempDir(), "upload-*.png")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer os.Remove(tempFile.Name())

			_, err = io.Copy(tempFile, resp.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file, err = os.Open(tempFile.Name())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer file.Close()

			header = &multipart.FileHeader{
				Filename: filepath.Base(tempFile.Name()),
				Size:     resp.ContentLength,
			}
		} else {
			defer file.Close()
		}

		req := azure.UploadImageParams{
			CompanyID:   companyID,
			CompanyName: companyName,
			LogoFile:    file,
			LogoHeader:  header,
		}

		err = cr.AzureStorage.UploadLogo(r.Context(), req)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(errors.ErrGeneratingAuthToken))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
