package goAddSvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func NewEndpoints(svc Service, logger log.Logger) Endpoints {
	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = makeAddEndpoint(svc)
		addEndpoint = endpointLoggingMiddleware(log.With(logger, "method", "Add"))(addEndpoint)
	}
	return Endpoints{addEndpoint: addEndpoint}
}

type Endpoints struct {
	addEndpoint endpoint.Endpoint
}

// Add implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (s Endpoints) Add(ctx context.Context, a, b RealNum) (RealNum, error) {
	resp, err := s.addEndpoint(ctx, addRequest{A: a, B: b})
	if err != nil {
		return 0, err
	}
	response := resp.(addResponse)
	return response.V, response.Err
}

func makeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addRequest)
		v, err := s.Add(ctx, req.A, req.B)
		return addResponse{V: v, Err: err}, nil
	}
}

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = addResponse{}
)

// addRequest collects the request parameters for the Add method.
type addRequest struct {
	A, B RealNum
}

// addResponse collects the response values for the Add method.
type addResponse struct {
	V   RealNum `json:"v"`
	Err error   `json:"-"` // should be intercepted by Failed/errorEncoder
}

// Failed implements endpoint.Failer.
func (r addResponse) Failed() error { return r.Err }
