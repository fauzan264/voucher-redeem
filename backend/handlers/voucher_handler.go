package handlers

import (
	"strings"

	"github.com/fauzan264/voucher-redeem/backend/constants"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/backend/helpers"
	"github.com/fauzan264/voucher-redeem/backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type voucherHandler struct {
	voucherService services.VoucherService
}

func NewVoucherHandler(voucherService services.VoucherService) *voucherHandler {
	return &voucherHandler{voucherService}
}

func (h *voucherHandler) CreateVoucher(c *fiber.Ctx) error {
	var request request.CreateVoucherRequest

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

	voucherResponse, err := h.voucherService.CreateVoucher(request)
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
		Data: voucherResponse,
	})
}

func (h *voucherHandler) GetVoucher(c *fiber.Ctx) error {
	var searchData request.SearchVoucher

	err := c.QueryParser(&searchData)
	if err != nil {
		if strings.Contains(err.Error(), "invalid UUID length") {
			return c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Status: false,
				Message: constants.FailedGetData,
				Errors: []string{"voucher_id must be a valid UUID"},
				Data: nil,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{"Invalid query parameters"},
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(searchData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	voucherResponse, err := h.voucherService.GetVoucher(searchData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: voucherResponse,
	})
}

func (h *voucherHandler) GetAllVoucherByBrand(c *fiber.Ctx) error {
	var searchData request.SearchVoucherByBrand

	err := c.QueryParser(&searchData)
	if err != nil {
		if strings.Contains(err.Error(), "invalid UUID length") {
			return c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Status: false,
				Message: constants.FailedGetData,
				Errors: []string{"brand_id must be a valid UUID"},
				Data: nil,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{"Invalid query parameters"},
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(searchData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	listVoucherResponse, err := h.voucherService.GetAllVoucherByBrand(searchData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: listVoucherResponse,
	})
}
