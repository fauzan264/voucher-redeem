package handlers

import (
	"github.com/fauzan264/voucher-redeem/constants"
	"github.com/fauzan264/voucher-redeem/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/helpers"
	"github.com/fauzan264/voucher-redeem/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	var request request.RegisterRequest

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

	registerUserResponse, err := h.authService.RegisterUser(request)
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
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: registerUserResponse,
	})
}

func (h *authHandler) LoginUser(c *fiber.Ctx) error {
	var request request.LoginRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	loginUserResponse, err := h.authService.LoginUser(request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
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
		Data: loginUserResponse,
	})
}