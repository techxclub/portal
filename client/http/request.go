package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/techx/portal/client/http/wrapper"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/logger"
)

func NewRequest(ctx context.Context, cmdName string) *Request {
	return &Request{
		header:     make(http.Header),
		queryParam: make(url.Values),
		ctx:        ctx,
		cmdName:    cmdName,
	}
}

type Request struct {
	ctx        context.Context
	method     string
	path       string
	body       interface{}
	host       string
	header     http.Header
	queryParam url.Values
	cmdName    string
}

func (r *Request) SetMethod(m string) *Request {
	r.method = m
	return r
}

func (r *Request) SetPath(format string, args ...interface{}) *Request {
	r.path = fmt.Sprintf(format, args...)
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.body = body
	return r
}

func (r *Request) SetHost(host string) *Request {
	r.host = host
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	for h, v := range headers {
		r.header.Set(h, v)
	}
	return r
}

func (r *Request) SetHeader(header, value string) *Request {
	r.header.Set(header, value)
	return r
}

func (r *Request) SetQueryParam(param, value string) *Request {
	r.queryParam.Set(param, value)
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for p, v := range params {
		r.SetQueryParam(p, v)
	}
	return r
}

func (r *Request) Build() (*http.Request, error) {
	uri := &url.URL{
		Scheme: "http",
		Host:   r.host,
		Path:   r.path,
	}

	uri.RawQuery = r.queryParam.Encode()

	body, err := r.bodyReader()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(r.ctx, r.method, uri.String(), body)
	if err != nil {
		return nil, err
	}

	for k, values := range r.header {
		for _, val := range values {
			req.Header.Add(k, val)
		}
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

func (r *Request) bodyReader() (io.ReadCloser, error) {
	if r.body == nil {
		return nil, nil
	}

	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(r.body); err != nil {
		return nil, err
	}
	return wrapper.BytesNopCloser(buffer.Bytes()), nil
}

func (r *Request) Send(doer Doer, success, failure interface{}) error {
	req, err := r.Build()
	if err != nil {
		return errors.NewUpstreamError(err, r.cmdName, 0)
	}
	return send(doer, req, r.cmdName, success, failure)
}

func send(doer Doer, r *http.Request, cmdName string, success, failure interface{}) error {
	logFields := map[string]interface{}{logger.ExternalServiceField: cmdName}
	resp, err := doer.Do(r)
	if resp != nil { // wrapper/do_hystrix.go returns both resp & err in case of 5xx
		defer func() {
			if err := drainAndClose(resp.Body); err != nil {
				logger.HTTPError(r, err, logFields)
			}
		}()
	}
	if err != nil {
		logger.HTTPError(r, err, logFields)
		statusCode := 0
		var targetError errors.HystrixError
		if errors.As(err, &targetError) {
			statusCode = targetError.GetStatusCode()
		}
		return errors.NewUpstreamError(err, cmdName, statusCode)
	}

	if err := decodeResponse(resp, success, failure); err != nil {
		logger.HTTPError(r, err, logFields)
		return errors.NewUpstreamError(err, cmdName, resp.StatusCode)
	}

	if err = checkResponseCode(resp.StatusCode); err != nil {
		logger.HTTPResponse(r, err, resp.StatusCode, success, logFields)
		return errors.NewUpstreamError(err, cmdName, resp.StatusCode)
	}

	logger.HTTPResponse(r, err, resp.StatusCode, success, logFields)
	return nil
}

func decodeResponse(resp *http.Response, successV, failureV interface{}) error {
	decoder := json.NewDecoder(resp.Body)

	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decoder.Decode(successV)
		}
	} else {
		if failureV != nil {
			return decoder.Decode(failureV)
		}
	}
	return nil
}

func drainAndClose(body io.ReadCloser) error {
	_, _ = io.Copy(io.Discard, body)
	return body.Close()
}

func checkResponseCode(code int) error {
	if code >= http.StatusBadRequest {
		return fmt.Errorf("received %d statuscode from client", code)
	}
	return nil
}
