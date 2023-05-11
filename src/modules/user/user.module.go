package user

import (
	"context"
	"template/src/generated/sqlc"

	"github.com/gofiber/fiber/v2"
)

type UserModule struct{}

func (module UserModule) Bundle(queries sqlc.Queries, context context.Context, App *fiber.App) {
	service := UserService{
		queries: queries,
		ctx:     context,
	}
	controller := UserController{
		service: service,
		app:     App,
	}
	controller.initRoutes()

}
