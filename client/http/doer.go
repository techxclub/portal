package http

import (
	"net/http"
	"time"

	"github.com/techx/portal/client/http/wrapper"
	"github.com/techx/portal/config"
)

type Doer = wrapper.Doer

type BasicParams struct {
	CmdName              string
	HTTPTimeout          time.Duration
	MaxConcurrentRequest int
	Transport            http.RoundTripper
}

func DefaultDoer(cmdName string, conf config.HTTPConfig) Doer {
	doer := BasicDoer(BasicParams{
		CmdName:              cmdName,
		HTTPTimeout:          time.Millisecond * time.Duration(conf.HTTPTimeout),
		MaxConcurrentRequest: conf.MaxConcurrentRequests,
		Transport:            conf.Transport,
	})

	return wrapper.DecorateDo(doer,
		wrapper.WithRetry(cmdName, conf.RetryCount, conf.RetryOnCodes),
		wrapper.WithHystrixMiddleware(cmdName, conf.HystrixConfig()),
	)
}

func BasicDoer(params BasicParams) Doer {
	return &http.Client{
		// additional buffer duration is added to trigger hystrix timeout first
		Timeout:   params.HTTPTimeout + time.Millisecond*20,
		Transport: getTransport(params.CmdName, params.Transport, params.MaxConcurrentRequest),
	}
}

func getTransport(_ string, transport http.RoundTripper, maxConcurrentRequest int) http.RoundTripper {
	if maxConcurrentRequest < 5 {
		maxConcurrentRequest = 5
	}

	// Buffer to handle cases like timeouts
	maxConcurrentRequestWithBuffer := int(1.2 * float64(maxConcurrentRequest))

	if transport == nil {
		transport = http.DefaultTransport
	}
	if t, ok := transport.(*http.Transport); ok {
		t = t.Clone()
		t.MaxIdleConnsPerHost = maxConcurrentRequestWithBuffer
		t.MaxIdleConns = maxConcurrentRequestWithBuffer
		t.MaxConnsPerHost = 0
		t.IdleConnTimeout = 3 * time.Minute
		transport = t
	}

	return transport
}
