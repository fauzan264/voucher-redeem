package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/voucher-redeem/backend/config"
	"github.com/fauzan264/voucher-redeem/backend/handlers"
	"github.com/fauzan264/voucher-redeem/backend/repositories"
	"github.com/fauzan264/voucher-redeem/backend/services"
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

	// services
	brandService := services.NewBrandService(brandRepository)

	// handlers
	brandHandler := handlers.NewBrandHandler(brandService)

	api := router.Group("/api/v1")

	// brand
	api.Post("/brand", brandHandler.CreateBrand)

	// voucher
	// api.Post("/voucher", voucherHandler.CreateVoucher)
	// api.Get("/voucher", voucherHandler.GetDetailVoucher)
	// api.Get("/voucher/brand", voucherHandler.GetAllVoucherByBrand)

	// transaction redemption
	// api.Post("/transaction/redemption", transactionHandler.CreateRedemption)
	// api.Get("/transaction/redemption", transactionHandler.GetDetailRedemption)
	
	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}