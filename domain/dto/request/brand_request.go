package request

type CreateBrandRequest struct {
	Name			string		`json:"name" validate:"required"`
	Description		string		`json:"description" validate:"required"`
}