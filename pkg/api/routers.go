package api

import (
	"github.com/dphantek/common/jwt"
	accountsApi "github.com/dphantek/common/pkg/api/accounts"
	authApi "github.com/dphantek/common/pkg/api/auth"
	cmsApi "github.com/dphantek/common/pkg/api/cms"
	mailApi "github.com/dphantek/common/pkg/api/mail"
	mediaApi "github.com/dphantek/common/pkg/api/media"
	siteApi "github.com/dphantek/common/pkg/api/site"
	wsApi "github.com/dphantek/common/pkg/api/ws"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlers(router fiber.Router, keyManager *jwt.KeyManager) {
	// router := app.Group("/api", auth.AuthMiddleware(keyManager, "/api"))

	siteApi.RegisterHandlers(router, keyManager)
	authApi.RegisterHandlers(router, keyManager)
	accountsApi.RegisterHandlers(router, keyManager)
	mailApi.RegisterHandlers(router)
	mediaApi.RegisterHandlers(router)
	cmsApi.RegisterHandlers(router)
	wsApi.RegisterHandlers(router)
}
