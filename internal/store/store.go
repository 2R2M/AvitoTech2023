package store

type Store interface {
	Segment() UserSegmentRepository
}
