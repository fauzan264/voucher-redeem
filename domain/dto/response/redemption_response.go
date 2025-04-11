package response

import (
	"time"

	"github.com/fauzan264/voucher-redeem/domain/models"
	"github.com/google/uuid"
)

type RedemptionResponse struct {
	ID					uuid.UUID					`json:"id"`
	Code 				string  					`json:"code"`
	TotalPoint			int							`json:"point_used"`
	RedeemedAt			time.Time					`json:"redeem_at"`

	Customer			UserResponse				`json:"customer"`
	RedemptionItems		[]RedemptionItemResponse	`json:"redemption_items"`
}

type RedemptionItemResponse struct {
	ID					uuid.UUID					`json:"id"`
	Quantity			int							`json:"quantity"`
	SubTotalPoint		int							`json:"sub_total_point"`

	VoucherResponse		VoucherResponse				`json:"voucher"`
}

func RedemptionItemResponseFormatter(redemptionItem models.RedemptionItem) RedemptionItemResponse {
	redemptionItemResponse := RedemptionItemResponse{
		ID: redemptionItem.ID,
		Quantity: redemptionItem.Quantity,
		SubTotalPoint: redemptionItem.SubTotalPoint,
	}

	redemptionItemResponse.VoucherResponse = VoucherResponseFormatter(redemptionItem.Voucher)

	return redemptionItemResponse
}

func ListRedemptionItemResponseFormatter(listRedemptionItem []models.RedemptionItem) []RedemptionItemResponse {
	var listRedemptionItemResponse []RedemptionItemResponse
	for _, redemption := range listRedemptionItem {
		redemptionItemResponseFormatter := RedemptionItemResponseFormatter(redemption)
		
		listRedemptionItemResponse = append(listRedemptionItemResponse, redemptionItemResponseFormatter)
	}

	return listRedemptionItemResponse
}

func RedemptionResponseFormatter(redemption models.Redemption) RedemptionResponse {
	redemptionResponse := RedemptionResponse{
		ID: redemption.ID,
		Code: redemption.Code,
		TotalPoint: redemption.TotalPoint,
		RedeemedAt: redemption.RedeemedAt,
	}

	redemptionResponse.Customer = UserResponse{
		ID: redemption.Customer.ID,
		Name: redemption.Customer.Name,
		Email: redemption.Customer.Email,
		TotalPoints: redemption.Customer.TotalPoints,
		IsAdmin: redemption.Customer.IsAdmin,
	}

	redemptionResponse.RedemptionItems = ListRedemptionItemResponseFormatter(redemption.RedemptionItems)

	return redemptionResponse
}