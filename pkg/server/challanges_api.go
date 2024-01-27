package server

import (
	"github.com/gofiber/fiber/v2"
	"time"

	"code.smartsheep.studio/hydrogen/passport/pkg/security"
	"code.smartsheep.studio/hydrogen/passport/pkg/services"
	"github.com/samber/lo"
)

func startChallenge(c *fiber.Ctx) error {
	var data struct {
		ID string `json:"id"`
	}

	if err := BindAndValidate(c, &data); err != nil {
		return err
	}

	user, err := services.LookupAccount(data.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	factors, err := services.LookupFactorsByUser(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	challenge, err := security.NewChallenge(user, factors, c.IP(), c.Get(fiber.HeaderUserAgent))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(fiber.Map{
		"display_name": user.Nick,
		"challenge":    challenge,
		"factors":      factors,
	})
}

func doChallenge(c *fiber.Ctx) error {
	var data struct {
		ChallengeID uint   `json:"challenge_id"`
		FactorID    uint   `json:"factor_id"`
		Secret      string `json:"secret"`
	}

	if err := BindAndValidate(c, &data); err != nil {
		return err
	}

	challenge, err := services.LookupChallengeWithFingerprint(data.ChallengeID, c.IP(), c.Get(fiber.HeaderUserAgent))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	factor, err := services.LookupFactor(data.FactorID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := security.DoChallenge(challenge, factor, data.Secret); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	challenge, err = services.LookupChallenge(data.ChallengeID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if challenge.Progress >= challenge.Requirements {
		session, err := security.GrantSession(challenge, []string{"*"}, nil, lo.ToPtr(time.Now()))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(fiber.Map{
			"is_finished": true,
			"challenge":   challenge,
			"session":     session,
		})
	}

	return c.JSON(fiber.Map{
		"is_finished": false,
		"challenge":   challenge,
		"session":     nil,
	})
}

func exchangeToken(c *fiber.Ctx) error {
	var data struct {
		Code      string `json:"code"`
		GrantType string `json:"grant_type"`
	}

	if err := BindAndValidate(c, &data); err != nil {
		return err
	}

	switch data.GrantType {
	case "authorization_code":
		access, refresh, err := security.ExchangeToken(data.Code)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(fiber.Map{
			"access_token":  access,
			"refresh_token": refresh,
		})
	case "refresh_token":
		access, refresh, err := security.RefreshToken(data.Code)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(fiber.Map{
			"access_token":  access,
			"refresh_token": refresh,
		})
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Unsupported exchange token type.")
	}
}