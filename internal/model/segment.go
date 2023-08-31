package model

type Segment struct {
	Slug    string
	Percent float64
}

func NewSegment(Slug string, Percent float64) *Segment {
	return &Segment{Slug: Slug, Percent: Percent}
}
