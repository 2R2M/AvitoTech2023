package model

type Segment struct {
	Slug string
}

func NewSegment(Slug string) *Segment {
	return &Segment{Slug: Slug}
}
