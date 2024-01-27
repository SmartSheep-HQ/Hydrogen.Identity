package server

import (
	"code.smartsheep.studio/hydrogen/passport/pkg/security"
	"code.smartsheep.studio/hydrogen/passport/pkg/services"
	"github.com/gofiber/fiber/v2"
)

func requestFactorToken(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("factorId", 0)

	factor, err := services.LookupFactor(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if sent, err := security.GetFactorCode(factor); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	} else if !sent {
		return c.SendStatus(fiber.StatusNoContent)
	} else {
		return c.SendStatus(fiber.StatusOK)
	}
}