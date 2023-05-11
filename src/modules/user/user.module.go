package user

import (
	"template/src/common"

	"github.com/gofiber/fiber/v2"
)

type UserModule common.Module

func (module UserModule) createController() common.IController {
	service := UserService{
		Imports: module.Imports,
	}
	controller := UserController{
		service: service,
		imports: module.Imports,
	}
	return controller
}

func (module UserModule) Bundle() fiber.Router {
	module.Controller = module.createController()
	return module.Controller.InitRoutes()

}
