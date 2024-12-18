package service

import (
	"encoding/json"
	"fmt"
	"git-functionality/internal/models"
	"git-functionality/internal/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VersionService struct {
	db          *gorm.DB
	userRepo    *repository.UserRepository
	versionRepo *repository.VersionRepository
}

func NewVersionService(
	db *gorm.DB,
	userRepo *repository.UserRepository,
	versionRepo *repository.VersionRepository,
) *VersionService {
	return &VersionService{
		db:          db,
		userRepo:    userRepo,
		versionRepo: versionRepo,
	}
}

func (s *VersionService) CommitChanges(
	userID uuid.UUID,
	changes map[string]interface{},
	commitMessage string,
) error {
	// Start a transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Validate user
		var user models.User
		if err := tx.First(&user, "id = ?", userID).Error; err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		// Create version record
		newVersion := &models.UserVersion{
			ID:            uuid.New(),
			UserID:        userID,
			ParentVersion: user.LatestVersionID,
			BranchName:    "main",
			CommitMessage: commitMessage,
			CreatedAt:     time.Now(),
		}

		// Serialize changes
		changesJSON, err := json.Marshal(changes)
		if err != nil {
			return err
		}
		newVersion.Changes = string(changesJSON)

		// Save the new version
		if err := tx.Create(newVersion).Error; err != nil {
			return err
		}

		// Prepare update fields
		updateFields := map[string]interface{}{
			"latest_version_id": newVersion.ID,
		}

		// Add changes to update fields
		for key, value := range changes {
			updateFields[key] = value
		}

		// Update user
		if err := tx.Model(&user).Updates(updateFields).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *VersionService) ValidateChanges(
	user *models.User,
	changes map[string]interface{},
) error {
	// Convert user to map for validation
	userMap := make(map[string]interface{})
	userBytes, _ := json.Marshal(user)
	json.Unmarshal(userBytes, &userMap)

	// Check if changes are valid
	for key := range changes {
		if _, exists := userMap[key]; !exists {
			return fmt.Errorf("invalid field: %s", key)
		}
	}

	return nil
}

func (s *VersionService) GetVersionHistory(
	userID uuid.UUID,
	branchName string,
) ([]models.UserVersion, error) {
	return s.versionRepo.GetVersionHistory(userID, branchName)
}

func (s *VersionService) RollbackToVersion(
	userID, versionID uuid.UUID,
) error {
	return s.versionRepo.RollbackToVersion(userID, versionID)
}
