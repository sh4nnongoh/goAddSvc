package goAddSvc

import (
	"context"

	"github.com/go-kit/kit/log"
)

// serviceMiddleware describes a service (as opposed to endpoint) middleware.
type serviceMiddleware func(Service) Service

// serviceLoggingMiddleware takes a logger as a dependency
// and returns a serviceMiddleware.
func serviceLoggingMiddleware(logger log.Logger) serviceMiddleware {
	return func(next Service) Service {
		return svcLoggingMiddleware{logger, next}
	}
}

type svcLoggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw svcLoggingMiddleware) Add(ctx context.Context, a, b float64) (v float64, err error) {
	defer func() {
		mw.logger.Log("method", "Add", "a", a, "b", b, "v", v, "err", err)
	}()
	return mw.next.Add(ctx, a, b)
}
