package common

import (
	"context"
	"template/src/generated/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IModule interface {
	Bundle(sqlc.Queries, context.Context, *fiber.App)
}
