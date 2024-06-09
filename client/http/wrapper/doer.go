package wrapper

import (
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type DoMiddleware func(doer Doer) Doer

func DecorateDo(doer Doer, middlewares ...DoMiddleware) Doer {
	for i := len(middlewares) - 1; i >= 0; i-- {
		doer = middlewares[i](doer)
	}

	return doer
}
