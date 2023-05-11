package src

import (
	"context"
	"template/src/common"
	"template/src/generated/sqlc"
	"template/src/modules/user"

	"github.com/gofiber/fiber/v2"
)

type AppModule struct {
	modules []common.IModule
	App     *fiber.App
	Ctx     context.Context
	Queries *sqlc.Queries
}

func (app AppModule) Bootstrap() {
	builder := app.CreateAppBuilder()
	builder.AddModule(user.UserModule{
		Imports: common.Imports{
			Env:     common.Env{},
			App:     app.App,
			Ctx:     app.Ctx,
			Queries: *app.Queries,
		},
	})
	builder.Build()
}
