package apicontext

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/techx/portal/constants"
)

const (
	keyRequestContext = "request-context"
)

type RequestContext struct {
	TraceID string
}

func NewRequestContextFromHTTP(r *http.Request) RequestContext {
	return RequestContext{
		TraceID: getRequestTraceID(r.Header),
	}
}

func NewContextWithRequestContext(ctx context.Context, reqCtx RequestContext) context.Context {
	return WithValue(ctx, keyRequestContext, reqCtx)
}

func RequestContextFromContext(ctx context.Context) RequestContext {
	c, ok := Value(ctx, keyRequestContext).(RequestContext)
	if !ok {
		return RequestContext{}
	}
	return c
}

func getRequestTraceID(header http.Header) string {
	for _, h := range []string{constants.HeaderXRequestTraceID} {
		if v := header.Get(h); v != "" {
			return v
		}
	}

	return uuid.NewV4().String()
}
