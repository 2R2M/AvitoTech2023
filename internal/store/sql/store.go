package sql

import (
	"avitoTech/internal/store"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db      *sqlx.DB
	segment *UserSegmentRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Segment() store.UserSegmentRepository {
	if s.segment != nil {
		return s.segment
	}
	return &UserSegmentRepository{store: s}
}
