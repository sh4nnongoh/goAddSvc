package addsvc

import "context"

type Service interface {
	Add(ctx context.Context, r []RealNum) (RealNum, error)
}

type RealNum float64
