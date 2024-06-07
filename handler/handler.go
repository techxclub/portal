package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/logger"
)

type ErrorDetails struct {
	Code            string `json:"code"`
	Message         string `json:"message"`
	MessageSeverity string `json:"message_severity"`
}

// swagger:model
type ErrorResponse struct {
	Errors []ErrorDetails `json:"errors"`
}

type validator interface {
	Validate() error
}

func Handler[RequestType validator, DomainType, ResponseType any](
	requestProcessor func(*http.Request) (*RequestType, error),
	processor func(context.Context, RequestType) (*DomainType, error),
	respProcessor func(context.Context, DomainType) (ResponseType, response.HTTPMetadata),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respBody, httpMetadata, err := process(r, requestProcessor, processor, respProcessor)
		if httpMetadata.Headers != nil {
			setHeaders(w, *httpMetadata.Headers)
		}

		if httpMetadata.Cookies != nil {
			http.SetCookie(w, httpMetadata.Cookies)
		}

		if err != nil {
			renderErrorResponse(r, w, err)
			return
		}
		writeJSON(w, http.StatusOK, respBody)
	}
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	body, err := json.Marshal(v)
	if err != nil {
		msg := fmt.Sprintf("Could not convert the given object to JSON: %v", err)
		writeText(w, http.StatusInternalServerError, msg)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func writeText(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(body))
}

func process[RequestType validator, DomainType, ResponseType any](
	r *http.Request,
	requestProcessor func(*http.Request) (*RequestType, error),
	processor func(context.Context, RequestType) (*DomainType, error),
	respProcessor func(context.Context, DomainType) (ResponseType, response.HTTPMetadata),
) (*ResponseType, response.HTTPMetadata, errors.ServiceError) {
	req, err := requestProcessor(r)
	if err != nil {
		return nil, response.HTTPMetadata{}, errors.BadRequestError(err)
	}

	if err := (*req).Validate(); err != nil {
		return nil, response.HTTPMetadata{}, errors.AsServiceError(err)
	}

	domainObj, err := processor(r.Context(), *req)
	if err != nil {
		return nil, response.HTTPMetadata{}, errors.AsServiceError(err)
	}

	respBody, metadata := respProcessor(r.Context(), *domainObj)
	return &respBody, metadata, nil
}

func setHeaders(w http.ResponseWriter, respHeaders http.Header) {
	for k, values := range respHeaders {
		for _, v := range values {
			if v != "" {
				w.Header().Add(k, v)
			}
		}
	}
}

func renderErrorResponse(r *http.Request, w http.ResponseWriter, serviceError errors.ServiceError) {
	logErrorResponse(r, serviceError)

	errResponse := customErrorResponse(r.Context(), serviceError)
	writeJSON(w, serviceError.GetResponseStatus(), errResponse)
}

func customErrorResponse(_ context.Context, serviceError errors.ServiceError) ErrorResponse {
	e := ErrorDetails{
		Code:            serviceError.GetCode(),
		Message:         serviceError.Error(),
		MessageSeverity: "error",
	}

	return ErrorResponse{
		Errors: []ErrorDetails{e},
	}
}

func logErrorResponse(r *http.Request, serviceError errors.ServiceError) {
	fields := map[string]interface{}{
		logger.ResponseStatusField: serviceError.GetResponseStatus(),
		"code":                     serviceError.GetCode(),
	}
	logger.HTTPError(r, serviceError, fields)
}
