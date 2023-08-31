package services

import (
	"avitoTech/internal/model"
	"avitoTech/internal/store"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

type UserSegmentService struct {
	userSegmentRepo store.UserSegmentRepository
}

func NewUserSegmentService(userSegmentRepo store.UserSegmentRepository) *UserSegmentService {
	return &UserSegmentService{
		userSegmentRepo: userSegmentRepo,
	}
}

func (s UserSegmentService) CreateUser(ctx context.Context, in *model.User) (*model.User, error) {
	user := model.NewUser(in.ID)
	newUser, err := s.userSegmentRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("service create user: %w", err)
	}
	return newUser, nil
}

func (s UserSegmentService) CreateSegment(ctx context.Context, in *model.Segment) (*model.Segment, error) {
	segment := model.NewSegment(in.Slug)
	newSegment, err := s.userSegmentRepo.CreateSegment(ctx, segment)
	if err != nil {
		return nil, fmt.Errorf("service create segment: %w", err)
	}
	return newSegment, nil
}

func (s UserSegmentService) DeleteSegment(ctx context.Context, slug string) error {

	errDelete := s.userSegmentRepo.DeleteSegment(ctx, slug)
	if errDelete != nil {
		return fmt.Errorf(errDelete.Error())
	}
	return nil
}

func (s UserSegmentService) AddUsersSegment(ctx context.Context, in *model.Operation) (*model.Operation, error) {
	operation := model.NewOperation(in.UserId, in.Segment, in.ExpiredAt)
	operations, err := s.userSegmentRepo.AddUsersSegment(ctx, operation)
	if err != nil {
		return nil, fmt.Errorf("service add segment for user: %w", err)
	}
	return operations, nil
}

func (s UserSegmentService) DeleteUsersSegment(ctx context.Context, in *model.Operation) (*model.Operation, error) {
	operation := model.NewOperation(in.UserId, in.Segment, in.ExpiredAt)
	operations, err := s.userSegmentRepo.DeleteUsersSegment(ctx, operation)
	if err != nil {
		return nil, fmt.Errorf("service delete segment for user: %w", err)
	}
	return operations, nil
}

func (s UserSegmentService) GetUsersSegments(ctx context.Context, id string) ([]*model.Segment, error) {
	return s.userSegmentRepo.GetUsersSegments(ctx, id)
}

func (s UserSegmentService) GetReport(ctx context.Context, month string, year string) ([]byte, error) {
	reports, errGetReports := s.userSegmentRepo.GetReport(ctx, month, year)
	if errGetReports != nil {
		return []byte{}, fmt.Errorf("repo get reports from db: %v", errGetReports.Error())
	}
	fileName := fmt.Sprintf("report_%s_%s.csv", year, month)
	file, err := os.Create(fileName)
	if err != nil {
		return []byte{}, fmt.Errorf("service create file error: %w", err)
	}

	var csvData bytes.Buffer
	writer := csv.NewWriter(&csvData)

	headers := []string{"User ID", "Slug", "Operation", "Created at"}
	writer.Write(headers)

	for _, report := range reports {
		row := []string{report.UserID, report.SegmentID, report.OpType, report.CreatedAt}
		writer.Write(row)
	}
	writer.Flush()
	file.Close()
	return csvData.Bytes(), nil
}
