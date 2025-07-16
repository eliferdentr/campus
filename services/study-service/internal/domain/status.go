package domain

// Status represents the status of an operation or entity.
type Status string

const (
	StatusPending   Status = "pending"
	StatusApproved  Status = "approved"
	StatusRejected  Status = "rejected"
	StatusActive    Status = "active"
	StatusInactive  Status = "inactive"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusCancelled Status = "cancelled"
)

// IsValid checks if the status is one of the predefined valid statuses.
func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusApproved, StatusRejected, StatusActive, StatusInactive, StatusCompleted, StatusFailed, StatusCancelled:
		return true
	}
	return false
}
