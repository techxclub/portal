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
	Origin   string
	Language string
	TraceID  string
	UserUUID string
}

func NewRequestContextFromHTTP(r *http.Request) RequestContext {
	return RequestContext{
		Origin:   r.Header.Get(constants.HeaderOrigin),
		Language: constants.DefaultLanguage,
		TraceID:  getRequestTraceID(r.Header),
		UserUUID: r.Header.Get(constants.HeaderXUserUUID),
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

func (r RequestContext) GetOrigin() string {
	return r.Origin
}

func (r RequestContext) GetLocale() string {
	s := strings.ReplaceAll(r.Language, "-", "_")
	s = strings.Split(s, "_")[0]
	s = strings.ToLower(s)
	if s == "" {
		return constants.DefaultLanguage
	}
	return s
}

func (r RequestContext) GetUserUUID() string {
	return r.UserUUID
}

func (r RequestContext) GetTraceID() string {
	return r.TraceID
}

func getRequestTraceID(header http.Header) string {
	for _, h := range []string{constants.HeaderXRequestTraceID} {
		if v := header.Get(h); v != "" {
			return v
		}
	}

	return utils.GetRandomUUID()
}
