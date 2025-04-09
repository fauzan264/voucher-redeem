package services

import (
	"errors"
	"time"

	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/fauzan264/voucher-redeem/backend/repositories"
	"github.com/google/uuid"
)

type voucherService struct {
	voucherRepository repositories.VoucherRepository
	brandRepository repositories.BrandRepository
}

type VoucherService interface {
	CreateVoucher(requestData request.CreateVoucherRequest) (response.VoucherResponse, error)
	GetVoucher(requestSearch request.SearchVoucher) (response.VoucherResponse, error)
	GetAllVoucherByBrand(requestSearch request.SearchVoucherByBrand) ([]response.VoucherResponse, error)
}

func NewVoucherService(
	voucherRepository repositories.VoucherRepository,
	brandRepository repositories.BrandRepository,
) *voucherService {
	return &voucherService{voucherRepository, brandRepository}
}

func (s *voucherService) CreateVoucher(requestData request.CreateVoucherRequest) (response.VoucherResponse, error) {
	brand, err := s.brandRepository.GetBrand(requestData.BrandID)
	if err != nil {
		return response.VoucherResponse{}, errors.New("Brand ID not found.")
	}
	
	voucherData := models.Voucher{
		ID: uuid.New(),
		BrandID: requestData.BrandID,
		Code: requestData.Code,
		Name: requestData.Name,
		Description: requestData.Description,
		CostInPoints: requestData.CostInPoints,
		Stock: requestData.Stock,
		ValidFrom: requestData.ValidFrom,
		ValidUntil: requestData.ValidUntil,
		CreatedAt: time.Now(),
	}
	
	voucher, err := s.voucherRepository.CreateVoucher(voucherData)
	if err != nil {
		return response.VoucherResponse{}, err
	}

	voucher.Brand = brand

	voucherResponse := response.VoucherResponseFormatter(voucher)

	return voucherResponse, nil
}

func (s *voucherService) GetVoucher(requestSearch request.SearchVoucher) (response.VoucherResponse, error) {
	voucher, err := s.voucherRepository.GetVoucher(requestSearch)
	if err != nil {
		return response.VoucherResponse{}, err
	}

	brand, _ := s.brandRepository.GetBrand(voucher.BrandID)
	voucher.Brand = brand

	voucherResponse := response.VoucherResponseFormatter(voucher)

	return voucherResponse, nil
}

func (s *voucherService) GetAllVoucherByBrand(requestSearch request.SearchVoucherByBrand) ([]response.VoucherResponse, error) {
	listVoucher, err := s.voucherRepository.GetVoucherByBrand(requestSearch)
	if err != nil {
		return []response.VoucherResponse{}, err
	}

	listVoucherResponse := response.ListVoucherResponseFormatter(listVoucher)

	return listVoucherResponse, nil
}