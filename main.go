package main

import (
	"template/src"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	appModule := src.AppModule{
		App: app,
	}
	appModule.Bootstrap()
	app.Listen(":3000")
}
