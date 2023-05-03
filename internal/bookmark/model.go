package bookmark

import (
	"time"

	"github.com/lib/pq"
)

type Bookmark struct {
	UserID    string
	JobPostID string
	CreatedAt time.Time
	AppliedAt pq.NullTime

	JobSlug       string
	JobTitle      string
	CompanyName   string
	JobExternalID string

	HasApplyRecord bool
}
