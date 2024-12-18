package models

import (
	"time"

	"github.com/google/uuid"
)

// UserVersion represents a version of user data
type UserVersion struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;"`                    // primary key
	UserID        uuid.UUID `gorm:"type:uuid;not null"`                        // foreign key
	ParentVersion uuid.UUID `gorm:"type:uuid;"`                                // points to the parent version
	BranchName    string    `gorm:"type:varchar(255);default:'main';not null"` // default to "main"
	CommitMessage string    `gorm:"type:varchar(255);not null"`                // commit message
	Changes       string    `gorm:"type:jsonb;not null"`                       // Store changes in JSON format
	CreatedAt     time.Time
}
