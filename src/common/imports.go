package common

import (
	"context"
	"template/src/generated/sqlc"

	"github.com/gofiber/fiber/v2"
)

type Imports struct {
	Env     Env
	App     *fiber.App
	Ctx     context.Context
	Queries sqlc.Queries
}
