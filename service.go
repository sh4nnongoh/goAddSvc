package goAddSvc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
)

type Service interface {
	Add(context.Context, RealNum, RealNum) (RealNum, error)
}

func NewService(logger log.Logger) Service {
	var svc Service
	{
		svc = service{}
		svc = serviceLoggingMiddleware(log.With(logger, "method", "Add"))(svc)
	}
	return svc
}

type service struct{}

func (s service) Add(_ context.Context, a, b RealNum) (RealNum, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

type RealNum float64

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

var (
	// ErrTwoZeroes is an arbitrary business rule for the Add method.
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	// ErrIntOverflow protects the Add method. We've decided that this error
	// indicates a misbehaving service and should count against e.g. circuit
	// breakers. So, we return it directly in endpoints, to illustrate the
	// difference. In a real service, this probably wouldn't be the case.
	ErrIntOverflow = errors.New("integer overflow")

	// ErrMaxSizeExceeded protects the Concat method.
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)
