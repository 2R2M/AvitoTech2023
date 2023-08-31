package store

import (
	"avitoTech/internal/model"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.33.0 --name=UserSegmentRepository

type UserSegmentRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CreateSegment(ctx context.Context, segment *model.Segment) (*model.Segment, error)
	DeleteSegment(ctx context.Context, slug string) error
	AddUsersSegment(ctx context.Context, operation *model.Operation) (*model.Operation, error)
	DeleteUsersSegment(ctx context.Context, operation *model.Operation) (*model.Operation, error)
	GetUsersSegments(ctx context.Context, id string) ([]*model.Segment, error)
	GetReport(ctx context.Context, month string, year string) ([]*model.Report, error)
}
