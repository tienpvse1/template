package user

import (
	"strconv"
	"template/src/common"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service UserService
	imports common.Imports
}

func (controller UserController) findById(c *fiber.Ctx) error {
	id, castErr := strconv.Atoi(c.Params("id"))
	if castErr != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	user, err := controller.service.findById(int32(id))
	if err != nil {
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(user)

}

func (controller UserController) InitRoutes() fiber.Router {
	return controller.imports.App.Route("/user", func(router fiber.Router) {
		router.Get("/:id", controller.findById)
	})
}
