package wrapper

import (
	"context"
	"io"
	"net/http"
	"slices"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func WithRetry(cmdName string, retryCount int, retryOnCodes []int) DoMiddleware {
	if retryCount <= 0 {
		return func(next Doer) Doer { return next }
	}

	return func(next Doer) Doer {
		return &retryMiddleware{cmdName, next, retryCount, retryOnCodes}
	}
}

type retryMiddleware struct {
	cmdName      string
	next         Doer
	retryCount   int
	retryOnCodes []int
}

func (s *retryMiddleware) Do(r *http.Request) (resp *http.Response, err error) {
	var body io.ReadSeekCloser
	body, err = s.retryableRequestBody(r)
	if err != nil {
		return
	}
	r.Body = body

	resp, err = s.next.Do(r)
	for i := 0; i < s.retryCount && s.isRetryable(resp, err); i++ {
		if body != nil {
			_, _ = body.Seek(0, 0)
		}

		if resp != nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}

		time.Sleep(time.Millisecond * 5)
		resp, err = s.next.Do(r)
	}

	return
}

func (s *retryMiddleware) retryableRequestBody(request *http.Request) (io.ReadSeekCloser, error) {
	if request.Body == nil {
		return nil, nil
	}
	if v, ok := request.Body.(*bytesNopCloser); ok {
		return v, nil
	}

	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return BytesNopCloser(b), nil
}

func (s *retryMiddleware) isRetryable(response *http.Response, err error) bool {
	switch err {
	case hystrix.ErrCircuitOpen, hystrix.ErrMaxConcurrency,
		context.DeadlineExceeded, context.Canceled:
		return false
	case nil:
		if response == nil {
			return false
		}
		statusCode := response.StatusCode
		return statusCode >= http.StatusInternalServerError ||
			slices.Contains(s.retryOnCodes, statusCode)
	default:
		switch e := err.(type) {
		case interface{ GetStatusCode() int }:
			statusCode := e.GetStatusCode()
			return statusCode >= http.StatusInternalServerError ||
				slices.Contains(s.retryOnCodes, statusCode)
		default:
			return true
		}
	}
}
