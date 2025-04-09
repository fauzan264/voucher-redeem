package repositories

import (
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"gorm.io/gorm"
)

type voucherRepository struct {
	db *gorm.DB
}

type VoucherRepository interface {
	CreateVoucher(voucher models.Voucher) (models.Voucher, error)
}

func NewVoucherRepository(db *gorm.DB) *voucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) CreateVoucher(voucher models.Voucher) (models.Voucher, error) {
	err := r.db.Create(&voucher).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}