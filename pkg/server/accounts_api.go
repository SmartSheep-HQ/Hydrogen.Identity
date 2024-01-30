package server

import (
	"code.smartsheep.studio/hydrogen/passport/pkg/database"
	"code.smartsheep.studio/hydrogen/passport/pkg/models"
	"code.smartsheep.studio/hydrogen/passport/pkg/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func getPrincipal(c *fiber.Ctx) error {
	user := c.Locals("principal").(models.Account)

	var data models.Account
	if err := database.C.
		Where(&models.Account{BaseModel: models.BaseModel{ID: user.ID}}).
		Preload("Profile").
		Preload("Contacts").
		Preload("Factors").
		Preload("Sessions").
		Preload("Challenges").
		First(&data).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(data)
}

func doRegister(c *fiber.Ctx) error {
	var data struct {
		Name       string `json:"name"`
		Nick       string `json:"nick"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		MagicToken string `json:"magic_token"`
	}

	if err := BindAndValidate(c, &data); err != nil {
		return err
	} else if viper.GetBool("use_registration_magic_token") && len(data.MagicToken) <= 0 {
		return fmt.Errorf("missing magic token in request")
	}

	if tk, err := services.ValidateMagicToken(data.MagicToken, models.RegistrationMagicToken); err != nil {
		return err
	} else {
		database.C.Delete(&tk)
	}

	if user, err := services.CreateAccount(
		data.Name,
		data.Nick,
		data.Email,
		data.Password,
	); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else {
		return c.JSON(user)
	}
}

func doRegisterConfirm(c *fiber.Ctx) error {
	var data struct {
		Code string `json:"code"`
	}

	if err := BindAndValidate(c, &data); err != nil {
		return err
	}

	if err := services.ConfirmAccount(data.Code); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
