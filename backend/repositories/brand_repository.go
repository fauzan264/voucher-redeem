package repositories

import (
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type brandRepository struct {
	db *gorm.DB
}

type BrandRepository interface {
	CreateBrand(brand models.Brand) (models.Brand, error)
	GetBrand(id uuid.UUID) (models.Brand, error)
}

func NewBrandRepository(db *gorm.DB) *brandRepository {
	return &brandRepository{db}
}

func (r *brandRepository) CreateBrand(brand models.Brand) (models.Brand, error) {
	err := r.db.Create(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}

func (r *brandRepository) GetBrand(id uuid.UUID) (models.Brand, error) {
	var brand models.Brand

	err := r.db.Where("id = ?", id).First(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}