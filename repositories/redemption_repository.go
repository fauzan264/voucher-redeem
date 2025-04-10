package repositories

import (
	"github.com/fauzan264/voucher-redeem/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type redemptionRepository struct {
	db *gorm.DB
}

type RedemptionRepository interface {
	CreateTransactionRedemption(redemption models.Redemption) (models.Redemption, error)
	GetTransactionRedemption(requestSearch request.SearchTransactionRedemption) (models.Redemption, error)
}

func NewRedemptionRepository(db *gorm.DB) *redemptionRepository {
	return &redemptionRepository{db}
}

func (r *redemptionRepository) CreateTransactionRedemption(redemption models.Redemption) (models.Redemption, error) {
	err := r.db.Create(&redemption).Error
	if err != nil {
		return redemption, err
	}

	return redemption, nil
}

func (r *redemptionRepository) GetTransactionRedemption(requestSearch request.SearchTransactionRedemption) (models.Redemption, error) {
	var redemption models.Redemption
	
	query := r.db.Model(&models.Redemption{})

	if requestSearch.TransactionID != uuid.Nil {
		query = query.Where("id = ?", requestSearch.TransactionID)
	}

	err := query.Preload("Customer").Preload("RedemptionItems.Voucher.Brand").First(&redemption).Error;
	if err != nil {
		return redemption, err
	}

	return redemption, nil
}
