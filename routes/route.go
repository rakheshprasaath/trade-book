package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakheshprasaath/trade-book.git/controller"
)

func Setup(app *fiber.App){
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/addAccount", controller.AddAccount)
	app.Get("/api/getAccounts", controller.GetAccounts)
	app.Get("/api/currentPositions/:accountKey", controller.GetCurrentPositionsByAccountKey)
}