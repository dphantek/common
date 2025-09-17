package auth

import (
	"strings"

	"github.com/dphantek/common/api"
	"github.com/dphantek/common/pkg/site"
	"github.com/dphantek/common/system"
	"github.com/gofiber/fiber/v2"
)

func APIKeyMiddleware(apikeyKey string, args ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header, prefix := "X-Api-Key", ""
		if len(args) > 0 {
			header = args[0]
		}
		if len(args) > 1 {
			prefix = args[1]
		}

		key := ctx.Get(header)
		if key != "" && prefix != "" {
			key, _ = strings.CutPrefix(key, prefix+" ")
		}

		if key == "" {
			return api.ErrorUnauthorizedResp(ctx, "API key is missing")
		}

		if apikey := site.Settings().CustomParams.GetString(apikeyKey); key != apikey {
			system.Logger.Error("settings key:", apikeyKey, "apikey:", apikey, "key:", key)
			return api.ErrorUnauthorizedResp(ctx, "Invalid API key")
		}

		return ctx.Next()
	}
}
