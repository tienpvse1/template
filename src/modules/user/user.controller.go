package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service UserService
	app     *fiber.App
}

func (controller UserController) findById(c *fiber.Ctx) error {
	// id, castErr := strconv.Atoi(c.Params("id"))
	// if castErr != nil {
	// 	return c.JSON(fiber.ErrBadRequest)
	// }
	// user, err := controller.service.findById(int32(id))
	// if err != nil {
	// 	return c.JSON(fiber.ErrNotFound)
	// }
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func (controller UserController) initRoutes() fiber.Router {
	return controller.app.Route("/user", func(router fiber.Router) {
		router.Get("/:id", controller.findById)
	})
}
