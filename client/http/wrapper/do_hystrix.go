package wrapper

import (
	"context"
	"net/http"
	"slices"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/techx/portal/errors"
)

func WithHystrixMiddleware(cmdName string, hystrixConf hystrix.CommandConfig, breakOnCodes ...int) DoMiddleware {
	return func(doer Doer) Doer {
		hystrix.ConfigureCommand(cmdName, hystrixConf)
		return &hystrixMiddleware{
			cmdName:      cmdName,
			next:         doer,
			breakOnCodes: breakOnCodes,
		}
	}
}

type hystrixMiddleware struct {
	cmdName      string
	next         Doer
	breakOnCodes []int
}

func (h *hystrixMiddleware) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	err := hystrix.DoC(req.Context(), h.cmdName, func(_ context.Context) (err error) {
		//nolint
		resp, err = h.next.Do(req)
		if err != nil {
			return err
		}

		if resp != nil && (slices.Contains(h.breakOnCodes, resp.StatusCode) || resp.StatusCode >= 500) {
			return errors.NewHystrixError(h.cmdName, resp.StatusCode)
		}

		return nil
	}, nil)

	return resp, err
}
