package handlers

import (
	"github.com/fauzan264/voucher-redeem/backend/constants"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/backend/helpers"
	"github.com/fauzan264/voucher-redeem/backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type brandHandler struct {
	brandService services.BrandService
}

func NewBrandHandler(brandService services.BrandService) *brandHandler {
	return &brandHandler{brandService}
}

func (h *brandHandler) CreateBrand(c *fiber.Ctx) error {
	var request request.CreateBrandRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	brandResponse, err := h.brandService.CreateBrand(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessInsertData,
		Errors: nil,
		Data: brandResponse,
	})
}