package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/gookit/slog"
	"github.com/jmoiron/sqlx"
)

func ManageDBUserSegment(ctx context.Context, db *sqlx.DB) {
	for {
		select {
		case <-ctx.Done():
			slog.Error("Data copying stopped.")
			return
		default:
			var operations []struct {
				ID        string     `db:"id"`
				UserID    string     `db:"user_id"`
				SegmentID string     `db:"segment_id"`
				OpType    string     `db:"operation_type"`
				CreatedAt *time.Time `db:"created_at"`
				ExpiredAt *time.Time `db:"expires_at"`
				Status    string     `db:"status"`
			}

			qAllOperations := `SELECT * FROM operation WHERE status = 'PENDING' ORDER BY created_at asc`
			err := db.SelectContext(ctx, &operations, qAllOperations)
			if err != nil {
				slog.Errorf("ManageDBUSerSegment: Read all operations: %v", err.Error())

			}
			for _, operation := range operations {

				qStatusUpdateUser := `UPDATE operation SET status = :status WHERE id=:id`
				type User struct {
					ID string `db:"id"`
				}
				user := User{}

				qUser := `SELECT * FROM "user" WHERE id = $1`
				errGetUser := db.GetContext(ctx, &user, qUser, operation.UserID)
				if errGetUser != nil {
					_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdateUser, map[string]any{
						"id":     operation.ID,
						"status": "ERROR",
					})
					if errUpdateStatus != nil {
						slog.Errorf("check user ID: %v", err.Error())
					}

				}
				type Segment struct {
					Slug string `db:"slug"`
				}
				qStatusUpdateSlug := `UPDATE operation SET status = :status WHERE id=:id`
				qSegment := `SELECT * FROM segment WHERE slug = $1`
				segment := Segment{}
				errGetSlug := db.GetContext(ctx, &segment, qSegment, operation.SegmentID)
				if errGetSlug != nil {
					_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdateSlug, map[string]any{
						"id":     operation.ID,
						"status": "ERROR",
					})
					if errUpdateStatus != nil {
						slog.Errorf("check slug: %v", err.Error())
					}
				}

				if errGetUser == nil && errGetSlug == nil {
					switch operation.OpType {
					case "ADD":
						qInsertUserSegment := `INSERT INTO usersegment (user_id, slug, expired_at) VALUES (:user_id, :slug, :expired_at)`
						_, errAdd := db.NamedExecContext(ctx, qInsertUserSegment, map[string]any{
							"user_id":    operation.UserID,
							"slug":       operation.SegmentID,
							"expired_at": operation.ExpiredAt,
						})
						qStatusUpdate := `UPDATE operation SET status = :status WHERE id=:id`
						if errAdd != nil {
							_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdate, map[string]any{
								"id":     operation.ID,
								"status": "ERROR",
							})
							if errUpdateStatus != nil {
								slog.Errorf("operation status update adding: %v", errUpdateStatus.Error())
							}
						} else {
							_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdate, map[string]any{
								"id":     operation.ID,
								"status": "SUCCESS",
							})
							if errUpdateStatus != nil {
								slog.Errorf("operation status update adding: %v", errUpdateStatus.Error())
							}
						}
					case "DELETE":
						qDeleteUserSegment := `DELETE FROM usersegment WHERE user_id=:user_id AND slug=:slug`
						res, errAddDelete := db.NamedExecContext(ctx, qDeleteUserSegment, map[string]any{
							"user_id": operation.UserID,
							"slug":    operation.SegmentID,
						})
						countRows, errAffected := res.RowsAffected()
						fmt.Println(countRows)
						qStatusUpdate := `UPDATE operation SET status = :status WHERE id=:id`
						if errAddDelete != nil || countRows < 1 || errAffected != nil {
							_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdate, map[string]any{
								"id":     operation.ID,
								"status": "ERROR",
							})
							if errUpdateStatus != nil {
								slog.Errorf("operation status update deleting: %v", errUpdateStatus.Error())
							}
						} else {
							_, errUpdateStatus := db.NamedExecContext(ctx, qStatusUpdate, map[string]any{
								"id":     operation.ID,
								"status": "SUCCESS",
							})
							if errUpdateStatus != nil {
								slog.Errorf("operation status update deleting: %v", errUpdateStatus.Error())
							}
						}

					}
				}
				select {
				case <-ctx.Done():
					slog.Error("Data copying stopped.")
					return
				case <-time.After(time.Microsecond):
				}
			}

		}
	}
}
