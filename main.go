package main

import (
	"github.com/NewChakrit/golang_clean_architecture/adapters"
	"github.com/NewChakrit/golang_clean_architecture/entities"
	"github.com/NewChakrit/golang_clean_architecture/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&entities.Order{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8080")
}
