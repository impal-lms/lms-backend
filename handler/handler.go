package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/repository/mock"
	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	userService services.UserService
}

func New() *Handler {
	userRepo := mock.NewUserMockRepository()

	userService := services.UserService{userRepo}
	return &Handler{
		userService,
	}
}

func (h Handler) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}

func (h Handler) Login(ctx echo.Context) error {
	b := ctx.Request().Body
	var req services.LoginRequest
	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		logrus.Error(err)
		return err
	}

	token, err := h.userService.Login(req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h Handler) CreateUser(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "token needed",
		})
	}

	requesterId, err := h.userService.ValidateToken(token)
	if err != nil {
		logrus.Error(err)
		return err
	}

	b := ctx.Request().Body
	var req services.CreateUserRequest
	err = json.NewDecoder(b).Decode(&req)
	if err != nil {
		logrus.Error(err)
		return err
	}

	userId, err := h.userService.CreateUser(requesterId, req)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"created": userId,
	})
}

func (h Handler) GetUserById(ctx echo.Context) error {
	p := ctx.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		logrus.Error(err)
		return err
	}

	user, err := h.userService.GetUserByID(int64(id))
	if err != nil {
		logrus.Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"user": *user,
	})
}
