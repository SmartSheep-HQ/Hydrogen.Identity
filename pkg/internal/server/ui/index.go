package ui

import (
	"fmt"

	"git.solsynth.dev/hydrogen/passport/pkg/services"
	"git.solsynth.dev/hydrogen/passport/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func MapUserInterface(A *fiber.App, authFunc utils.AuthFunc) {
	authCheckWare := func(c *fiber.Ctx) error {
		var token string
		if cookie := c.Cookies(services.CookieAccessKey); len(cookie) > 0 {
			token = cookie
		}

		c.Locals("token", token)

		if err := authFunc(c); err != nil {
			uri := c.Request().URI().FullURI()
			return c.Redirect(fmt.Sprintf("/sign-in?redirect_uri=%s", string(uri)))
		} else {
			return c.Next()
		}
	}

	pages := A.Group("/").Name("Pages")

	pages.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/users/me")
	})

	pages.Get("/sign-up", signupPage)
	pages.Get("/sign-in", signinPage)
	pages.Get("/mfa", mfaRequestPage)
	pages.Get("/mfa/apply", mfaApplyPage)
	pages.Get("/authorize", authCheckWare, authorizePage)

	pages.Post("/sign-up", signupAction)
	pages.Post("/sign-in", signinAction)
	pages.Post("/mfa", mfaRequestAction)
	pages.Post("/mfa/apply", mfaApplyAction)
	pages.Post("/authorize", authCheckWare, authorizeAction)

	pages.Get("/users/me", authCheckWare, selfUserinfoPage)
}