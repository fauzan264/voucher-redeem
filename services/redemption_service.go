package services

import (
	"errors"
	"time"

	"github.com/fauzan264/voucher-redeem/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/domain/models"
	"github.com/fauzan264/voucher-redeem/helpers"
	"github.com/fauzan264/voucher-redeem/repositories"
	"github.com/google/uuid"
)

type redemptionService struct {
	redemptionRepository repositories.RedemptionRepository
	voucherRepository repositories.VoucherRepository
	brandRepository repositories.BrandRepository
	userRepository repositories.UserRepository
}

type RedemptionService interface {
	CreateTransactionRedemption(requestData request.CreateTransactionRedemptionRequest) (response.RedemptionResponse, error)
	GetTransactionRedemption(requestSearch request.SearchTransactionRedemption) (response.RedemptionResponse, error)
}

func NewRedemptionService(
	redemptionRepository repositories.RedemptionRepository,
	voucherRepository repositories.VoucherRepository,
	brandRepository repositories.BrandRepository,
	userRepository repositories.UserRepository,
) *redemptionService {
	return &redemptionService{
		redemptionRepository,
		voucherRepository,
		brandRepository,
		userRepository}
}

func (s *redemptionService) CreateTransactionRedemption(requestData request.CreateTransactionRedemptionRequest) (response.RedemptionResponse, error) {
	customer, err := s.userRepository.GetUserByID(requestData.CustomerID)
	if err != nil {
		return response.RedemptionResponse{}, errors.New("User Not Found")
	}
	
	var totalPoint int
	var redemptionItemsData []models.RedemptionItem
	for _, redemptionItem := range requestData.TransactionItemsRedemptionRequest {
		voucher, err := s.voucherRepository.GetVoucher(request.SearchVoucher{VoucherID: redemptionItem.VoucherID})
		if err != nil {
			return response.RedemptionResponse{}, errors.New("Voucher Not Found")
		}

		brand, _ := s.brandRepository.GetBrand(voucher.BrandID)

		subTotalPoint := redemptionItem.Quantity * voucher.CostInPoints

		redemptionItemData := models.RedemptionItem{
			ID: uuid.New(),
			VoucherID: redemptionItem.VoucherID,
			Quantity: redemptionItem.Quantity,
			SubTotalPoint: subTotalPoint,
		}

		redemptionItemData.Voucher = voucher
		redemptionItemData.Voucher.Brand = brand
		totalPoint += subTotalPoint

		redemptionItemsData = append(redemptionItemsData, redemptionItemData)
	}

	customer.TotalPoints += totalPoint

	redemptionData := models.Redemption{
		ID: uuid.New(),
		CustomerID: customer.ID,
		Code: helpers.GenerateTransactionRedeemCode(),
		TotalPoint: totalPoint,
		RedeemedAt: time.Now(),
	}

	redemptionData.Customer = customer
	redemptionData.RedemptionItems = redemptionItemsData

	redemption, err := s.redemptionRepository.CreateTransactionRedemption(redemptionData)
	if err != nil {
		return response.RedemptionResponse{}, err
	}

	for i, updateVoucher := range redemptionItemsData {
		updateVoucher.Voucher.Stock -= updateVoucher.Quantity

		updatedVoucher, err := s.voucherRepository.UpdateVoucher(updateVoucher.Voucher)
		if err != nil {
			return response.RedemptionResponse{}, err
		}

		redemptionItemsData[i].Voucher = updatedVoucher
	}

	_, err = s.userRepository.UpdateUser(customer)
	if err != nil {
		return response.RedemptionResponse{}, err
	}

	redemptionResponse := response.RedemptionResponseFormatter(redemption)

	return redemptionResponse, nil
}

func (s *redemptionService) GetTransactionRedemption(requestSearch request.SearchTransactionRedemption) (response.RedemptionResponse, error) {
	redemption, err := s.redemptionRepository.GetTransactionRedemption(requestSearch)
	if err != nil {
		return response.RedemptionResponse{}, err
	}

	redemptionResponse := response.RedemptionResponseFormatter(redemption)

	return redemptionResponse, nil
}