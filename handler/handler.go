package handler

import (
	"context"
	"net/http"

	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/composers"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/utils"
)

type validator interface {
	Validate() error
}

func Handler[RequestType validator, DomainType, ResponseType any](
	requestProcessor func(*http.Request) (*RequestType, error),
	processor func(context.Context, RequestType) (*DomainType, error),
	respProcessor func(context.Context, DomainType) (ResponseType, composers.HTTPMetadata),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respBody, httpMetadata, err := process(r, requestProcessor, processor, respProcessor)
		if err != nil {
			response.RenderErrorResponse(r, w, err)
			return
		}

		if httpMetadata.Headers != nil {
			setHeaders(w, *httpMetadata.Headers)
		}

		if httpMetadata.Cookies != nil {
			http.SetCookie(w, httpMetadata.Cookies)
		}
		utils.WriteJSON(w, http.StatusOK, respBody)
	}
}

func process[RequestType validator, DomainType, ResponseType any](
	r *http.Request,
	requestProcessor func(*http.Request) (*RequestType, error),
	processor func(context.Context, RequestType) (*DomainType, error),
	respProcessor func(context.Context, DomainType) (ResponseType, composers.HTTPMetadata),
) (*ResponseType, composers.HTTPMetadata, errors.ServiceError) {
	req, err := requestProcessor(r)
	if err != nil {
		return nil, composers.HTTPMetadata{}, errors.BadRequestError(err)
	}

	if err := (*req).Validate(); err != nil {
		return nil, composers.HTTPMetadata{}, errors.AsServiceError(err)
	}

	domainObj, err := processor(r.Context(), *req)
	if err != nil {
		return nil, composers.HTTPMetadata{}, errors.AsServiceError(err)
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
