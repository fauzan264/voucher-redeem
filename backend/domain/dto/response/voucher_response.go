package response

import (
	"time"

	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/google/uuid"
)

type VoucherResponse struct {
	ID				uuid.UUID	`json:"id"`
    // BrandID      uuid.UUID 	`json:"brand_id"`
    Code         	string    		`json:"code"`
    Name         	string    		`json:"name"`
    Description  	string    		`json:"description"`
    CostInPoints 	int       		`json:"cost_in_points"`
    Stock        	int       		`json:"stock"`
    ValidFrom    	time.Time 		`json:"valid_from"`
    ValidUntil   	time.Time 		`json:"valid_until"`
	Brand			BrandResponse	`json:"brand"`
}

func VoucherResponseFormatter(voucher models.Voucher) VoucherResponse {
	voucherResponse := VoucherResponse{
		ID: voucher.ID,
		Code: voucher.Code,
		Name: voucher.Name,
		Description: voucher.Description,
		CostInPoints: voucher.CostInPoints,
		Stock: voucher.Stock,
		ValidFrom: voucher.ValidFrom,
		ValidUntil: voucher.ValidUntil,
	}

	voucherResponse.Brand = BrandResponse{
		ID: voucher.Brand.ID,
		Name: voucher.Brand.Name,
		Description: voucher.Brand.Description,
	}

	return voucherResponse
}

func ListVoucherResponseFormatter(listVoucher []models.Voucher) []VoucherResponse {
	var listVoucherResponse []VoucherResponse
	for _, voucher := range listVoucher {
		voucherResponseFormatter := VoucherResponseFormatter(voucher)
		
		listVoucherResponse = append(listVoucherResponse, voucherResponseFormatter)
	}

	return listVoucherResponse
}