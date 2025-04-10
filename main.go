package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/voucher-redeem/config"
	"github.com/fauzan264/voucher-redeem/handlers"
	"github.com/fauzan264/voucher-redeem/middleware"
	"github.com/fauzan264/voucher-redeem/repositories"
	"github.com/fauzan264/voucher-redeem/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitDatabase()

	router := fiber.New()
	router.Use(cors.New())

	// repositories
	brandRepository := repositories.NewBrandRepository(db)
	voucherRepository := repositories.NewVoucherRepository(db)
	userRepository := repositories.NewUserRepository(db)
	redemptionRepository := repositories.NewRedemptionRepository(db)

	// services
	brandService := services.NewBrandService(brandRepository)
	voucherService := services.NewVoucherService(voucherRepository, brandRepository)
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)
	redemptionService := services.NewRedemptionService(
		redemptionRepository,
		voucherRepository,
		brandRepository,
		userRepository,
	)

	// handlers
	brandHandler := handlers.NewBrandHandler(brandService)
	voucherHandler := handlers.NewVoucherHandler(voucherService)
	authHandler := handlers.NewAuthHandler(authService)
	redemptionHandler := handlers.NewRedemptionHandler(redemptionService)

	// middleware
	authMiddleware := middleware.AuthMiddleware(userService)

	// grouping
	api := router.Group("/api/v1")

	// auth
	api.Post("/auth/register", authHandler.RegisterUser)
	api.Post("/auth/login", authHandler.LoginUser)

	// brand
	api.Post("/brand", authMiddleware, brandHandler.CreateBrand)

	// voucher
	api.Post("/voucher", authMiddleware, voucherHandler.CreateVoucher)
	api.Get("/voucher", authMiddleware, voucherHandler.GetVoucher)
	api.Get("/voucher/brand", authMiddleware, voucherHandler.GetAllVoucherByBrand)

	// transaction redemption
	api.Post("/transaction/redemption", authMiddleware, redemptionHandler.CreateRedemption)
	api.Get("/transaction/redemption", authMiddleware, redemptionHandler.GetDetailRedemption)
	
	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}