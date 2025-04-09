package middleware

import (
	"strings"

	"github.com/fauzan264/voucher-redeem/backend/constants"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/backend/services"
	"github.com/fauzan264/voucher-redeem/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaderToken := c.Get("Authorization")

		if !strings.Contains(authHeaderToken, "Bearer") {
			return sendUnauthorizedResponse(c)
			
		}

		tokenString := ""
		arrayToken := strings.Split(authHeaderToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		jwtService := utils.NewJWTService()
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return sendUnauthorizedResponse(c)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return sendUnauthorizedResponse(c)
		}

		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return sendUnauthorizedResponse(c)
			
		}

		requestUser := request.GetUser{
			ID: userID,
		}

		user, err := userService.GetUserByID(requestUser)
		if err != nil {
			return sendUnauthorizedResponse(c)
		}

		c.Locals("authUser", user)

		return c.Next()
	}
}

func sendUnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
		Status: false,
		Message: constants.Unauthorized,
		Errors: []string{constants.ErrUnauthorized.Error()},
		Data: nil,
	})
}
