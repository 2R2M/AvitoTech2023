package sql

import (
	"avitoTech/internal/model"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserSegmentRepository struct {
	store *Store
}

func (s UserSegmentRepository) AddUsersSegment(ctx context.Context, operation *model.Operation) (*model.Operation, error) {

	timeRes := sql.NullTime{}
	expiredAtTime, errTime := time.Parse("2006-01-02 15:04:05", operation.ExpiredAt)
	if errTime != nil {
		timeRes.Valid = false
	}
	timeRes.Time = expiredAtTime
	qUserSegment := `INSERT INTO operation(id, user_id, segment_id, operation_type, expires_at, status) VALUES (:id, :user_id, :segment_id, :operation_type, :expires_at, :status)`
	for _, slug := range operation.Segment {
		_, createSQLErr := s.store.db.NamedExecContext(ctx, qUserSegment, map[string]any{
			"id":             uuid.New().String(),
			"user_id":        operation.UserId,
			"segment_id":     slug,
			"operation_type": "ADD",
			"expires_at":     timeRes,
			"status":         "PENDING",
		})
		if createSQLErr != nil {
			return nil, fmt.Errorf("repo create operation: %w", createSQLErr)
		}

	}
	return &model.Operation{
		ID:        operation.ID,
		UserId:    operation.UserId,
		Segment:   operation.Segment,
		ExpiredAt: operation.ExpiredAt,
	}, nil
}

func (s UserSegmentRepository) DeleteUsersSegment(ctx context.Context, operation *model.Operation) (*model.Operation, error) {

	qUserSegment := `INSERT INTO operation(id, user_id, segment_id, operation_type, status) VALUES (:id, :user_id, :segment_id, :operation_type, :status)`
	for _, slug := range operation.Segment {
		_, createSQLErr := s.store.db.NamedExecContext(ctx, qUserSegment, map[string]any{
			"id":             uuid.New().String(),
			"user_id":        operation.UserId,
			"segment_id":     slug,
			"operation_type": "DELETE",
			"status":         "PENDING",
		})
		if createSQLErr != nil {
			return nil, fmt.Errorf("repo delete operation: %w", createSQLErr)
		}
	}
	return &model.Operation{
		ID:      operation.ID,
		UserId:  operation.UserId,
		Segment: operation.Segment,
	}, nil
}

func (s UserSegmentRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	q := `INSERT INTO "user"(id) VALUES (:id)`

	_, createSQLErr := s.store.db.NamedExecContext(ctx, q, map[string]any{
		"id": user.ID,
	})
	if createSQLErr != nil {
		return nil, fmt.Errorf("repo create user: %w", createSQLErr)
	}

	return &model.User{
		ID: user.ID,
	}, nil
}

func (s UserSegmentRepository) CreateSegment(ctx context.Context, segment *model.Segment) (*model.Segment, error) {
	q := `INSERT INTO segment(slug, percent) VALUES (:slug, :percent)`

	_, createSQLErr := s.store.db.NamedExecContext(ctx, q, map[string]any{
		"slug":    segment.Slug,
		"percent": segment.Percent,
	})
	if createSQLErr != nil {
		return nil, fmt.Errorf("repo create segment: %w", createSQLErr)
	}

	return &model.Segment{
		Slug: segment.Slug,
	}, nil
}

func (s UserSegmentRepository) DeleteSegment(ctx context.Context, slug string) error {
	var users []struct {
		ID string `db:"user_id"`
	}

	q := `SELECT user_id FROM usersegment WHERE slug = $1;`
	err := s.store.db.SelectContext(ctx, &users, q, slug)
	if err != nil {
		return fmt.Errorf("repo get users from slug %s: %w", slug, err)
	}
	qUserSegment := `INSERT INTO operation(id, user_id, segment_id, operation_type, status) VALUES (:id, :user_id, :segment_id, :operation_type, :status)`
	for _, user := range users {
		_, createSQLErr := s.store.db.NamedExecContext(ctx, qUserSegment, map[string]any{
			"id":             uuid.New().String(),
			"user_id":        user.ID,
			"segment_id":     slug,
			"operation_type": "DELETE",
			"status":         "PENDING",
		})
		if createSQLErr != nil {
			return fmt.Errorf("repo delete slug operation: %w", createSQLErr)
		}
	}

	qDeleteSegment := `DELETE FROM segment WHERE slug=:slug`
	_, err = s.store.db.NamedExecContext(ctx, qDeleteSegment, map[string]any{
		"slug": slug,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("repo delete segment: %w", err)
	}

	return nil
}

func (s UserSegmentRepository) GetUsersSegments(ctx context.Context, id string) ([]*model.Segment, error) {
	var segments []struct {
		Slug string `db:"slug"`
	}
	q := `SELECT slug FROM usersegment WHERE user_id = $1 AND ( expired_at > NOW() OR expired_at IS NULL);`
	err := s.store.db.SelectContext(ctx, &segments, q, id)
	if err != nil {
		return nil, fmt.Errorf("repo get users segment slugs: %w", err)
	}
	listSegments := make([]*model.Segment, 0)
	for _, segment := range segments {
		listSegments = append(listSegments, &model.Segment{
			Slug: segment.Slug,
		})
	}
	return listSegments, nil
}

func (s UserSegmentRepository) GetReport(ctx context.Context, month string, year string) ([]*model.Report, error) {
	qReport := `SELECT user_id, segment_id, operation_type, created_at FROM operation 
                                                       WHERE status ='SUCCESS' AND  EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2`

	var Report []struct {
		UserID    string `db:"user_id"`
		SegmentId string `db:"segment_id"`
		OpType    string `db:"operation_type"`
		CreatedAt string `db:"created_at"`
	}
	err := s.store.db.SelectContext(ctx, &Report, qReport, year, month)
	if err != nil {
		return []*model.Report{}, fmt.Errorf("repo get CSV Report: %w", err)
	}

	listReports := make([]*model.Report, 0)
	for _, report := range Report {
		listReports = append(listReports, &model.Report{
			UserID:    report.UserID,
			SegmentID: report.SegmentId,
			OpType:    report.OpType,
			CreatedAt: report.CreatedAt,
		})
	}
	return listReports, nil

}
