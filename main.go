package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rakheshprasaath/trade-book.git/database"
	"github.com/rakheshprasaath/trade-book.git/routes"
)
func getPublicIP() string {
	resp, err := http.Get("https://api64.ipify.org")
	if err != nil {
		return "Error fetching public IP"
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response"
	}

	return string(ip)
}

func main() {
	fmt.Println("Public IP:", getPublicIP())

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
