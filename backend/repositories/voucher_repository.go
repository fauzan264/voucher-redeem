package repositories

import (
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type voucherRepository struct {
	db *gorm.DB
}

type VoucherRepository interface {
	CreateVoucher(voucher models.Voucher) (models.Voucher, error)
	GetVoucher(requestSearch request.SearchVoucher) (models.Voucher, error)
	GetVoucherByBrand(requestSearch request.SearchVoucherByBrand) ([]models.Voucher, error)
	UpdateVoucher(voucher models.Voucher) (models.Voucher, error)
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

func (r *voucherRepository) GetVoucher(requestSearch request.SearchVoucher) (models.Voucher, error) {
	var voucher models.Voucher
	
	query := r.db.Model(&models.Voucher{})

	if requestSearch.VoucherID != uuid.Nil {
		query = query.Where("id = ?", requestSearch.VoucherID)
	}

	err := query.First(&voucher).Error;
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (r *voucherRepository) GetVoucherByBrand(requestSearch request.SearchVoucherByBrand) ([]models.Voucher, error) {
	var listVoucher []models.Voucher

	query := r.db.Model(&models.Voucher{})

	if requestSearch.BrandID != uuid.Nil {
		query = query.Where("brand_id = ?", requestSearch.BrandID)
	}

	err := query.Preload("Brand").Find(&listVoucher).Error
	if err != nil {
		return listVoucher, err
	}

	return listVoucher, nil
}

func (r *voucherRepository) UpdateVoucher(voucher models.Voucher) (models.Voucher, error) {
	err := r.db.Save(&voucher).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}