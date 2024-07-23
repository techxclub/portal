package logger

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/errors"
)

var loggableRequestHeaders = []string{
	constants.HeaderXUserUUID,
	constants.HeaderXForwardedFor,
	constants.HeaderXRequestTraceID,
}

func getLogFieldForInterface(key string, val interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: val,
	}
}

func HTTPRequestLogger(r *http.Request) zerolog.Logger {
	logger := log.With().
		Str(RequestMethodField, r.Method).
		Str(RequestURLField, r.URL.RequestURI()).
		Str(RequestProxyField, r.RemoteAddr).
		Str(RequestTraceID, apicontext.RequestContextFromContext(r.Context()).TraceID)

	headers := map[string]string{}
	for _, header := range loggableRequestHeaders {
		if headerValue := r.Header.Get(header); headerValue != "" {
			headers[header] = headerValue
		}
	}
	logger = logger.Fields(getLogFieldForInterface(RequestHeaders, headers))
	return logger.Logger()
}

func HTTPError(request *http.Request, err error, fields map[string]interface{}) {
	l := httpBaseLogger(request, fields).With().Err(err).Logger()
	switch {
	case errors.Is(err, context.Canceled):
		l.Debug().Msg("received error")
	default:
		l.Error().Msg("received error")
	}
}

func HTTPResponse(request *http.Request, err error, status int, body interface{}, fields map[string]interface{}) {
	logger := httpBaseLogger(request, fields).With().
		Int(ResponseStatusField, status).Fields(getLogFieldForInterface(ResponseField, body)).Logger()
	switch {
	case err != nil:
		logger = logger.With().Err(err).Logger()
		logger.Error().Msg("received error")
	case status >= 400:
		logger.Error().Msg("received error")
	default:
		logger.Debug().Msg("received response")
	}
}

func httpBaseLogger(request *http.Request, fields map[string]interface{}) zerolog.Logger {
	return HTTPRequestLogger(request).With().Fields(fields).Logger()
}
