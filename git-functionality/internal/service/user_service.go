package service

import (
	"git-functionality/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	user := &models.User{
		ID:              uuid.New(),
		Name:            name,
		Email:           email,
		LatestVersionID: uuid.Nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       gorm.DeletedAt{},
	}

	result := s.db.Create(user)
	return user, result.Error
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	return users, result.Error
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, "id = ?", userID).Error
	return &user, err
}

func (s *UserService) UpdateUser(userID uuid.UUID, updates map[string]interface{}) error {
	return s.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(updates).Error
}
