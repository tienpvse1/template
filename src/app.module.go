package src

import (
	"context"
	"database/sql"
	"template/src/common"
	"template/src/generated/sqlc"
	"template/src/modules/user"

	"github.com/gofiber/fiber/v2"
)

type AppModule struct {
	modules []common.IModule
	App     *fiber.App
}

func (app AppModule) Bootstrap() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(db)
	builder := app.createAppBuilder()
	builder.addModule(user.UserModule{})
	builder.done(*queries, ctx)

}

func (app *AppModule) addModule(module any) *AppModule {
	app.modules = append(app.modules, module.(common.IModule))
	return app
}
func (app *AppModule) done(queries sqlc.Queries, ctx context.Context) {
	for _, m := range app.modules {
		m.Bundle(queries, ctx, app.App)
	}
}

func (app AppModule) createAppBuilder() *AppModule {
	return &app
}
