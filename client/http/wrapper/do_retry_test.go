package wrapper_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/stretchr/testify/assert"
	"github.com/techx/portal/client/http/wrapper"
	"github.com/techx/portal/errors"
)

func TestRetry(t *testing.T) {
	t.Parallel()
	clientTimeout := 200 * time.Millisecond
	reqBodyStr := `{"hello":"world"}`
	respBodyStr := `{"foo":"bar"}`
	type response struct {
		body         string
		statusCode   int
		err          error
		responseTime time.Duration
	}
	tt := []struct {
		name           string
		requestBody    io.Reader
		requestBodyStr string
		retryOnCode    []int
		middlewares    []wrapper.DoMiddleware
		responses      []response

		expected response
	}{
		{
			name:           "2xx",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			responses: []response{{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			}},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name: "2xx without request body",
			responses: []response{{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			}},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "4xx",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			responses: []response{{
				body:       respBodyStr,
				statusCode: http.StatusBadRequest,
			}},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:           "5xx followed by Success",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			responses: []response{
				{
					statusCode: http.StatusInternalServerError,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "4xx(retry on code) followed by Success",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			retryOnCode:    []int{http.StatusBadRequest},
			responses: []response{
				{
					statusCode: http.StatusBadRequest,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "5xx followed by Success with hystix middleware",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			middlewares: []wrapper.DoMiddleware{wrapper.WithHystrixMiddleware(
				"test",
				hystrix.CommandConfig{},
			)},
			responses: []response{
				{
					statusCode: http.StatusInternalServerError,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "4xx(retry on code) followed by Success with hystrix middleware",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			middlewares: []wrapper.DoMiddleware{wrapper.WithHystrixMiddleware(
				"test",
				hystrix.CommandConfig{},
			)},
			retryOnCode: []int{http.StatusBadRequest},
			responses: []response{
				{
					statusCode: http.StatusBadRequest,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "5xx followed by Success when body is already bytes nop closer",
			requestBody:    wrapper.BytesNopCloser([]byte(reqBodyStr)),
			requestBodyStr: reqBodyStr,
			responses: []response{
				{
					statusCode: http.StatusInternalServerError,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "hystrix circuit open",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			middlewares: []wrapper.DoMiddleware{wrapper.WithHystrixMiddleware(
				"test_open_circuit",
				hystrix.CommandConfig{},
			)},
			expected: response{
				err: hystrix.ErrCircuitOpen,
			},
		},
		{
			name:           "Timeout followed by Success",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			responses: []response{
				{
					body:         "test",
					responseTime: clientTimeout + (2 * time.Millisecond),
					statusCode:   http.StatusOK,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusOK,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusOK,
			},
		},
		{
			name:           "Exceeds max retry",
			requestBody:    strings.NewReader(reqBodyStr),
			requestBodyStr: reqBodyStr,
			responses: []response{
				{
					body:         "test",
					responseTime: clientTimeout + (2 * time.Millisecond),
					statusCode:   http.StatusOK,
				}, {
					body:       "test",
					statusCode: http.StatusInternalServerError,
				}, {
					body:       respBodyStr,
					statusCode: http.StatusServiceUnavailable,
				},
			},
			expected: response{
				body:       respBodyStr,
				statusCode: http.StatusServiceUnavailable,
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var respCount int
			mu := sync.Mutex{}
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "HEADER", r.Header.Get("TEST"))
				if tc.requestBodyStr != "" {
					b, err := io.ReadAll(r.Body)
					assert.NoError(t, err)
					assert.Equal(t, tc.requestBodyStr, string(b))
				}

				mu.Lock()
				defer mu.Unlock()
				if respCount > len(tc.responses)-1 {
					rw.WriteHeader(http.StatusInternalServerError)
				}
				resp := tc.responses[respCount]
				respCount++
				time.Sleep(resp.responseTime)
				rw.WriteHeader(resp.statusCode)
				if resp.body != "" {
					_, _ = rw.Write([]byte(resp.body))
				}
			}))
			defer server.Close()
			c := server.Client()
			c.Timeout = clientTimeout

			middleware := append([]wrapper.DoMiddleware{wrapper.WithRetry("test", 2, tc.retryOnCode)}, tc.middlewares...)
			client := wrapper.DecorateDo(c, middleware...)

			openHystrixCircuit("test_open_circuit")
			req, err := http.NewRequest(http.MethodPost, server.URL, tc.requestBody)
			assert.NoError(t, err)
			req.Header.Set("TEST", "HEADER")
			resp, err := client.Do(req)
			assert.Equal(t, tc.expected.err, err)
			if err == nil {
				defer resp.Body.Close()
				assert.Equal(t, tc.expected.statusCode, resp.StatusCode)
				var body string
				if resp.Body != nil {
					b, err := io.ReadAll(resp.Body)
					assert.NoError(t, err)
					body = string(b)
				}

				assert.Equal(t, tc.expected.body, body)
			}
			assert.Equal(t, len(tc.responses), respCount)
		})
	}
}

func openHystrixCircuit(cmdName string) {
	var err error
	for err != hystrix.ErrCircuitOpen {
		err = hystrix.Do(cmdName, func() error {
			return errors.New("test")
		}, nil)
	}
}
