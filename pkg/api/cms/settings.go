package cmsApi

import (
	"github.com/dphantek/common/api"
	"github.com/dphantek/common/pkg/cms"
	"github.com/dphantek/common/types"
	"github.com/gofiber/fiber/v2"
)

func getSettings(ctx *fiber.Ctx) error {
	return api.SuccessResp(ctx, cms.Settings())
}

func updateSettings(ctx *fiber.Ctx) error {
	params := &cms.CmsParams{}
	// need to reset custom_params
	// or, it will be merged with the old params
	params.CustomParams = &types.Params{}
	if err := ctx.BodyParser(params); err != nil {
		return api.ErrorBadRequestResp(ctx, err.Error())
	}
	err := cms.UpdateSettings(params)
	return api.SuccessResp(ctx, err == nil)
}
