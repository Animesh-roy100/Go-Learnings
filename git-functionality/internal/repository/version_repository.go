package repository

import (
	"encoding/json"
	"git-functionality/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VersionRepository struct {
	db *gorm.DB
}

func NewVersionRepository(db *gorm.DB) *VersionRepository {
	return &VersionRepository{db: db}
}

func (r *VersionRepository) CreateVersion(
	userID uuid.UUID,
	changes map[string]interface{},
	commitMessage string,
	branchName string,
) (*models.UserVersion, error) {
	// Serialize changes
	changesJSON, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	// Get the latest version ID
	var user models.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	// Create new version
	newVersion := &models.UserVersion{
		ID:            uuid.New(),
		UserID:        userID,
		ParentVersion: user.LatestVersionID,
		BranchName:    branchName,
		CommitMessage: commitMessage,
		Changes:       string(changesJSON),
		CreatedAt:     time.Now(),
	}

	// Use transaction to ensure atomicity
	err = r.db.Transaction(func(tx *gorm.DB) error {
		// Create the new version
		if err := tx.Create(newVersion).Error; err != nil {
			return err
		}

		// Update user's latest version
		return tx.Model(&user).
			Select("latest_version_id").
			Update("latest_version_id", newVersion.ID).Error
	})

	return newVersion, err
}

func (r *VersionRepository) GetVersionHistory(
	userID uuid.UUID,
	branchName string,
) ([]models.UserVersion, error) {
	var versions []models.UserVersion
	err := r.db.Where("user_id = ? AND branch_name = ?", userID, branchName).
		Order("created_at DESC").
		Find(&versions).Error

	return versions, err
}

func (r *VersionRepository) RollbackToVersion(
	userID, versionID uuid.UUID,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find the specific version
		var version models.UserVersion
		if err := tx.First(&version,
			"id = ? AND user_id = ?", versionID, userID).Error; err != nil {
			return err
		}

		// Parse changes
		var changes map[string]interface{}
		if err := json.Unmarshal([]byte(version.Changes), &changes); err != nil {
			return err
		}

		// Update user with the version's changes
		return tx.Model(&models.User{}).
			Where("id = ?", userID).
			Updates(changes).Error
	})
}
