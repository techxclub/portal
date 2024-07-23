package response

import (
	"context"
	"net/http"

	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/i18n"
	"github.com/techx/portal/logger"
	"github.com/techx/portal/utils"
)

// swagger:model
type ErrorResponse struct {
	Code     string `json:"code"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

func RenderErrorResponse(r *http.Request, w http.ResponseWriter, serviceError errors.ServiceError) {
	InstrumentErrorResponse(r, serviceError)

	errResponse := customErrorResponse(r.Context(), serviceError)
	if traceID := apicontext.RequestContextFromContext(r.Context()).TraceID; traceID != "" {
		w.Header().Set(constants.HeaderXRequestTraceID, traceID)
	}
	utils.WriteJSON(w, serviceError.GetResponseStatus(), errResponse)
}

func customErrorResponse(ctx context.Context, serviceError errors.ServiceError) ErrorResponse {
	return ErrorResponse{
		Code:     serviceError.GetCode(),
		Title:    i18n.Title(ctx, serviceError.GetI18nKey(), serviceError.GetI18nValues()),
		Message:  i18n.Message(ctx, serviceError.GetI18nKey(), serviceError.GetI18nValues()),
		Severity: "error",
	}
}

func InstrumentErrorResponse(r *http.Request, serviceError errors.ServiceError) {
	fields := map[string]interface{}{
		logger.ResponseStatusField: serviceError.GetResponseStatus(),
		"code":                     serviceError.GetCode(),
	}

	logger.HTTPError(r, serviceError, fields)
}
