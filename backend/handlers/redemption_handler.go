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

type redemptionHandler struct {
	redemptionService services.RedemptionService
}

func NewRedemptionHandler(redemptionService services.RedemptionService) *redemptionHandler {
	return &redemptionHandler{redemptionService}
}


func (h *redemptionHandler) CreateRedemption(c *fiber.Ctx) error {
	var request request.CreateTransactionRedemptionRequest
	
	authUser := c.Locals("authUser")
	if authUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	request.CustomerID = user.ID

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

	redemptionResponse, err := h.redemptionService.CreateTransactionRedemption(request)
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
		Data: redemptionResponse,
	})
}

func (h *redemptionHandler) GetDetailRedemption(c *fiber.Ctx) error {
	var searchData request.SearchTransactionRedemption

	err := c.QueryParser(&searchData)
	if err != nil {
		if strings.Contains(err.Error(), "invalid UUID length") {
			return c.Status(fiber.StatusBadRequest).JSON(response.Response{
				Status: false,
				Message: constants.FailedGetData,
				Errors: []string{"id must be a valid UUID"},
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

	redemptionResponse, err := h.redemptionService.GetTransactionRedemption(searchData)
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
		Data: redemptionResponse,
	})
}