package response

import (
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/google/uuid"
)


type BrandResponse struct {
	ID				uuid.UUID	`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
}

func BrandResponseFormatter(brand models.Brand) BrandResponse {
	return BrandResponse{
		ID : brand.ID,
		Name : brand.Name,
		Description : brand.Description,
	}
}