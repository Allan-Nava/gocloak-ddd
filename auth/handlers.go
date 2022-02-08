package auth

/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *AuthService
}

// CreateLiveClipping godoc
// @Summary      CreateLiveClipping
// @Description  CreateLiveClipping
// @Tags         livecut
// @Accept json
// @Param request body liveClippingRestreamerApiRequest true "request"
// @Produce      json
// @Success      200  {object} createLiveClippingResponse
// @Router       /livecut/live_clipping [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var requestBody LoginRequest
	//
	return c.Status(http.StatusOK).JSON(requestBody)
}
