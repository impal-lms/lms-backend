package handler

import (
	"encoding/json"
	"net/http"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
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

func (h *Handler) Register(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var Req domain.RegisterRequest
	err := json.NewDecoder(body).Decode(&Req)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	User := domain.User{
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
	}

	User, code, err := h.Services.CreateUser(User)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = domain.UserResponse{
		ID:    User.ID,
		Name:  User.Name,
		Email: User.Email,
		Role:  User.Role,
	}
	response.Status = code

	return ctx.JSON(code, response)
}
