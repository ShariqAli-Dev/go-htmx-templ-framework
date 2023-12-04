package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/shareed2k/goth_fiber"
)

func GetUserFromSession(c *fiber.Ctx) (goth.User, error) {
	providerName, err := goth_fiber.GetProviderName(c)
	if err != nil {
		return goth.User{}, err
	}
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}
	value, err := goth_fiber.GetFromSession(providerName, c)
	if err != nil {
		return goth.User{}, err
	}
	session, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}
	user, err := provider.FetchUser(session)
	if err != nil {
		return goth.User{}, err
	}
	return user, nil
}
