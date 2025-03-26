package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakheshprasaath/trade-book.git/controller"
)

func Setup(app *fiber.App){
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/webhook", controller.WebSocketHandler())

}