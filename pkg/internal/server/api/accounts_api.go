package api

import (
	"fmt"
	"strconv"
	"time"

	"git.solsynth.dev/hydrogen/passport/pkg/internal/server/exts"

	"git.solsynth.dev/hydrogen/passport/pkg/internal/database"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/models"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/services"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

func lookupAccount(c *fiber.Ctx) error {
	probe := c.Query("probe")
	if len(probe) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "lookup probe is required")
	}

	user, err := services.LookupAccount(probe)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}

func searchAccount(c *fiber.Ctx) error {
	probe := c.Query("probe")
	if len(probe) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "search probe is required")
	}

	users, err := services.SearchAccount(probe)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}

func getUserinfo(c *fiber.Ctx) error {
	if err := exts.EnsureAuthenticated(c); err != nil {
		return err
	}
	user := c.Locals("user").(models.Account)

	var data models.Account
	if err := database.C.
		Where(&models.Account{BaseModel: models.BaseModel{ID: user.ID}}).
		Preload("Profile").
		Preload("Contacts").
		Preload("Badges").
		First(&data).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else {
		data.PermNodes = c.Locals("permissions").(map[string]any)
	}

	var resp fiber.Map
	raw, _ := jsoniter.Marshal(data)
	jsoniter.Unmarshal(raw, &resp)

	resp["sub"] = strconv.Itoa(int(data.ID))
	resp["family_name"] = data.Profile.FirstName
	resp["given_name"] = data.Profile.LastName
	resp["name"] = data.Name
	resp["email"] = data.GetPrimaryEmail().Content
	resp["preferred_username"] = data.Nick

	if data.Avatar != nil {
		resp["picture"] = *data.GetAvatar()
	}

	return c.JSON(resp)
}

func getEvents(c *fiber.Ctx) error {
	if err := exts.EnsureAuthenticated(c); err != nil {
		return err
	}
	user := c.Locals("user").(models.Account)
	take := c.QueryInt("take", 0)
	offset := c.QueryInt("offset", 0)

	var count int64
	var events []models.ActionEvent
	if err := database.C.
		Where(&models.ActionEvent{AccountID: user.ID}).
		Model(&models.ActionEvent{}).
		Count(&count).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := database.C.
		Order("created_at desc").
		Where(&models.ActionEvent{AccountID: user.ID}).
		Limit(take).
		Offset(offset).
		Find(&events).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"count": count,
		"data":  events,
	})
}

func editUserinfo(c *fiber.Ctx) error {
	if err := exts.EnsureAuthenticated(c); err != nil {
		return err
	}
	user := c.Locals("user").(models.Account)

	var data struct {
		Nick        string    `json:"nick" validate:"required,min=2,max=24"`
		Description string    `json:"description"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		Birthday    time.Time `json:"birthday"`
	}

	if err := exts.BindAndValidate(c, &data); err != nil {
		return err
	}

	var account models.Account
	if err := database.C.
		Where(&models.Account{BaseModel: models.BaseModel{ID: user.ID}}).
		Preload("Profile").
		First(&account).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	account.Nick = data.Nick
	account.Description = data.Description
	account.Profile.FirstName = data.FirstName
	account.Profile.LastName = data.LastName
	account.Profile.Birthday = &data.Birthday

	if err := database.C.Save(&account).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if err := database.C.Save(&account.Profile).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	services.InvalidAuthCacheWithUser(account.ID)

	return c.SendStatus(fiber.StatusOK)
}

func doRegister(c *fiber.Ctx) error {
	var data struct {
		Name       string `json:"name" validate:"required,lowercase,alphanum,min=2,max=16"`
		Nick       string `json:"nick" validate:"required,min=2,max=24"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required,min=4,max=32"`
		MagicToken string `json:"magic_token"`
	}

	if err := exts.BindAndValidate(c, &data); err != nil {
		return err
	} else if viper.GetBool("use_registration_magic_token") && len(data.MagicToken) <= 0 {
		return fmt.Errorf("missing magic token in request")
	} else if viper.GetBool("use_registration_magic_token") {
		if tk, err := services.ValidateMagicToken(data.MagicToken, models.RegistrationMagicToken); err != nil {
			return err
		} else {
			database.C.Delete(&tk)
		}
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
		Code string `json:"code" validate:"required"`
	}

	if err := exts.BindAndValidate(c, &data); err != nil {
		return err
	}

	if err := services.ConfirmAccount(data.Code); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
