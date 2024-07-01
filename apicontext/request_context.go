package apicontext

import (
	"context"
	"net/http"
	"strings"

	"github.com/techx/portal/constants"
	"github.com/techx/portal/utils"
)

const (
	keyRequestContext = "request-context"
)

type RequestContext struct {
	Language string
	TraceID  string
	UserID   string
}

func NewRequestContextFromHTTP(r *http.Request) RequestContext {
	return RequestContext{
		Language: constants.DefaultLanguage,
		TraceID:  getRequestTraceID(r.Header),
		UserID:   r.Header.Get(constants.HeaderXUserID),
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

func (rctx RequestContext) GetLocale() string {
	s := strings.ReplaceAll(rctx.Language, "-", "_")
	s = strings.Split(s, "_")[0]
	s = strings.ToLower(s)
	if s == "" {
		return constants.DefaultLanguage
	}
	return s
}

func getRequestTraceID(header http.Header) string {
	for _, h := range []string{constants.HeaderXRequestTraceID} {
		if v := header.Get(h); v != "" {
			return v
		}
	}

	return utils.GetRandomUUID()
}

func (rctx RequestContext) GetUserID() string {
	return rctx.UserID
}
