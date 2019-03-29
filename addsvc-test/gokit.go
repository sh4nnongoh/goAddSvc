package addsvc

import "context"
import "encoding/json"
import "errors"
import "net/http"
import "github.com/go-kit/kit/endpoint"
import httptransport "github.com/go-kit/kit/transport/http"

type RealNum float64
type Service struct {
}

func (s Service) Add(ctx context.Context, r []RealNum) (RealNum, error) {
	panic(errors.New("not implemented"))
}

type AddRequest struct {
	R []RealNum
}
type AddResponse struct {
	R   RealNum
	Err error
}

func MakeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		R, err := s.Add(ctx, req.R)
		return AddResponse{R: R, Err: err}, nil
	}
}

type Endpoints struct {
	Add endpoint.Endpoint
}

func NewHTTPHandler(endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/add", httptransport.NewServer(endpoints.Add, DecodeAddRequest, EncodeAddResponse))
	return m
}
func DecodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeAddResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
