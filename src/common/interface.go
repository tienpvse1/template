package common

import "github.com/gofiber/fiber/v2"

type IModule interface {
	Bundle() fiber.Router
}

type IController interface {
	InitRoutes() fiber.Router
}
