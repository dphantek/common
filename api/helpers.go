package api

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResp(c *fiber.Ctx, data interface{}, meta ...ApiResponseMeta) error {
	resp := ApiResponse{
		Success: true,
		Data:    data,
	}
	if len(meta) > 0 {
		resp.Meta = &meta[0]
	}
	return c.Status(fiber.StatusOK).JSON(&resp)
}

func ErrorResp(c *fiber.Ctx, err ApiError, meta ...ApiResponseMeta) error {
	resp := ApiResponse{
		Success: false,
		Error:   &err,
	}
	if len(meta) > 0 {
		resp.Meta = &meta[0]
	}
	code := fiber.StatusBadRequest
	if err.Code != 0 {
		code = err.Code
	}
	return c.Status(code).JSON(&resp)
}

func ErrorCodeResp(c *fiber.Ctx, code int, message ...string) error {
	msg := "API Error"
	if len(message) > 0 {
		msg = message[0]
	}
	return ErrorResp(c, ApiError{
		Code:    code,
		Message: msg,
	})
}

func ErrorNotFoundResp(c *fiber.Ctx, message ...string) error {
	return ErrorCodeResp(c, fiber.StatusNotFound, message...)
}

func ErrorUnauthorizedResp(c *fiber.Ctx, message ...string) error {
	return ErrorCodeResp(c, fiber.StatusUnauthorized, message...)
}

func ErrorBadRequestResp(c *fiber.Ctx, message ...string) error {
	return ErrorCodeResp(c, fiber.StatusBadRequest, message...)
}

func ErrorInternalServerErrorResp(c *fiber.Ctx, message ...string) error {
	return ErrorCodeResp(c, fiber.StatusInternalServerError, message...)
}
