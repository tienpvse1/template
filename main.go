package main

import (
	"context"
	"fmt"
	"template/src"

	"template/src/common"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	env := common.Env{}
	env.TryLoadEnv()
	app := fiber.New()
	queries := common.InitConnectionAndGetQueries(env)
	appModule := src.AppModule{
		App:     app,
		Ctx:     context.Background(),
		Queries: queries,
	}
	appModule.Bootstrap()
	app.Listen(fmt.Sprintf(":%s", env.Get("GO_PORT")))
}
