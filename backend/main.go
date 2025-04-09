package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/voucher-redeem/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()
	_ = config.InitDatabase()

	router := fiber.New()
	router.Use(cors.New())

	// repositories

	// services

	// handlers

	_ = router.Group("/api/v1")
	
	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}