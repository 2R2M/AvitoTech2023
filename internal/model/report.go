package model

type Report struct {
	UserID    string
	SegmentID string
	OpType    string
	CreatedAt string
}

func NewReport(userID string, segmentID string, opType string, createdAt string) *Report {
	return &Report{
		UserID: userID, SegmentID: segmentID, OpType: opType, CreatedAt: createdAt,
	}
}
