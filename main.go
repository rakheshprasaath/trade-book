package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rakheshprasaath/trade-book.git/database"
	"github.com/rakheshprasaath/trade-book.git/routes"
)

func main() {
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	routes.Setup(app)
	app.Listen(":" + port)
}
