package repositories

import (
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(id uuid.UUID) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}