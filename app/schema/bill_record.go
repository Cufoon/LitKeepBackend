package schema

import "time"

type BillRecord struct {
	ID     uint
	UserID string
	KindID string
	Type   int
	Value  float64
	Time   time.Time
	Mark   string
}

type BillRecordQuery struct {
	UserID    string
	KindID    string
	StartTime time.Time
	EndTime   time.Time
}

type BillRecordPageQuery struct {
	UserID string
	Page   int
}
