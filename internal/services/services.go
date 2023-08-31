package services

import (
	"context"
	"errors"
)

type contextKey string

// Services — struct for services abstraction.
type Services struct {
	UserSegmentService UserSegmentService
}

// servicesKey - unique key for extraction of services from context.
const servicesKey contextKey = "userSegment.services"

var ErrMissingServices = errors.New("missing services")

func New(userSegmentService UserSegmentService) Services {
	return Services{UserSegmentService: userSegmentService}
}

func Set(ctx context.Context, ss Services) context.Context {
	return context.WithValue(ctx, servicesKey, ss)
}

func Get(ctx context.Context) (Services, error) {
	ss, ok := ctx.Value(servicesKey).(Services)
	if !ok {
		return Services{}, ErrMissingServices
	}

	return ss, nil
}

func Must(ctx context.Context) Services {
	ss, err := Get(ctx)
	if err != nil {
		panic(err)
	}

	return ss
}
