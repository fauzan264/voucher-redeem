package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateVoucherRequest struct {
    BrandID      	uuid.UUID	`json:"brand_id" validate:"required"`
    Code         	string   	`json:"code" validate:"required"`
    Name         	string   	`json:"name" validate:"required"`
    Description  	string   	`json:"description" validate:"required"`
    CostInPoints 	int      	`json:"cost_in_points" validate:"required"`
    Stock        	int      	`json:"stock" validate:"required"`
    ValidFrom    	time.Time	`json:"valid_from" validate:"required"`
    ValidUntil   	time.Time	`json:"valid_until" validate:"required"`
}

type SearchVoucher struct {
	VoucherID		uuid.UUID 	`query:"id" validate:"required"`
}

type SearchVoucherByBrand struct {
	BrandID			uuid.UUID 	`query:"id" validate:"required"`
}