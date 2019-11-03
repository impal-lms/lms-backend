package handler

import (
	"encoding/json"
	"net/http"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo"
)

func (h *Handler) Login(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var loginReq domain.LoginRequest
	err := json.NewDecoder(body).Decode(&loginReq)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	loginRes, err := h.Services.Login(loginReq)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	response.Data = loginRes
	response.Status = http.StatusOK
	return ctx.JSON(http.StatusOK, response)
}
