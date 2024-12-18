package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the latest version of user information
type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name            string    `gorm:"type:varchar(255);not null"`
	Email           string    `gorm:"type:varchar(255);unique;not null"`
	LatestVersionID uuid.UUID `gorm:"type:uuid;not null"` // Links to the latest version
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
