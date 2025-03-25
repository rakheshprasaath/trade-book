package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakheshprasaath/trade-book.git/controller"
)

func Setup(app *fiber.App){
	app.Post("/api/Register", controller.Register)

}