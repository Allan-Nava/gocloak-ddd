package auth

/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */

import (
	"net/http"

	"github.com/Allan-Nava/gocloak-ddd/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *AuthService
}

// Login godoc
// @Summary      Login
// @Description  Login
// @Tags         livecut
// @Accept json
// @Param request body LoginRequest true "request"
// @Produce      json
// @Success      200  {object} AuthResponse
// @Router       /livecut/live_clipping [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var requestBody LoginRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	login, err := h.Service.Login(requestBody.Username, requestBody.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&utils.ApiError{Message: err.Error()})
	}
	//
	return c.Status(http.StatusOK).JSON(&login)
}
