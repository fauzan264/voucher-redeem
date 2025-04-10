package request

import (
	"github.com/google/uuid"
)

type SearchTransactionRedemption struct {
	TransactionID		uuid.UUID 	`query:"id" validate:"required,uuid4"`
}

type CreateTransactionRedemptionRequest struct {
	CustomerID							uuid.UUID					`json:"customer_id" validate:"required,uuid4"`

	TransactionItemsRedemptionRequest	[]RedemptionItemRequest		`json:"redemption_items" validate:"required,dive"`
}

type RedemptionItemRequest struct {
	VoucherID			uuid.UUID					`json:"voucher_id" validate:"required,uuid4"`
	Quantity			int							`json:"quantity"  validate:"required"`
}