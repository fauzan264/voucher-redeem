package services

import (
	"time"

	"github.com/fauzan264/voucher-redeem/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/domain/models"
	"github.com/fauzan264/voucher-redeem/repositories"
	"github.com/google/uuid"
)

type brandService struct {
	brandRepository repositories.BrandRepository
}

type BrandService interface {
	CreateBrand(requestData request.CreateBrandRequest) (response.BrandResponse, error)
}

func NewBrandService(brandRepository repositories.BrandRepository) *brandService {
	return &brandService{brandRepository}
}

func (s *brandService) CreateBrand(requestData request.CreateBrandRequest) (response.BrandResponse, error) {
	brandData := models.Brand{
		ID: uuid.New(),
		Name: requestData.Name,
		Description: requestData.Description,
		CreatedAt: time.Now(),
	}

	brand, err := s.brandRepository.CreateBrand(brandData)
	if err != nil {
		return response.BrandResponse{}, err
	}

	brandResponse := response.BrandResponseFormatter(brand)

	return brandResponse, nil
}