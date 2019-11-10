package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllUser(ctx echo.Context) error {
	var response Response

	Users, err := h.Services.GetAllUser()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	UsersResponse := make([]domain.UserResponse, 0)
	for _, val := range Users {
		UsersResponse = append(UsersResponse, domain.UserResponse{
			ID:    val.ID,
			Name:  val.Name,
			Email: val.Email,
			Role:  val.Role,
		})
	}

	response.Data = UsersResponse
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateUser(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.CreateUserRequest
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	User := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
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

func (h *Handler) GetUserByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	User, err := h.Services.GetUserByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = domain.UserResponse{
		ID:    User.ID,
		Name:  User.Name,
		Email: User.Email,
		Role:  User.Role,
	}
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateUser(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.UpdateUserRequest
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	User := domain.User{}
	User.ID = id

	if request.Name != "" {
		User.Name = request.Name
	}

	if request.Email != "" {
		User.Email = request.Email
	}

	User, err = h.Services.UpdateUser(User)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = domain.UserResponse{
		ID:    User.ID,
		Name:  User.Name,
		Email: User.Email,
		Role:  User.Role,
	}
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteUserById(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeleteUserById(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) ChangePassword(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.ChangePasswordRequest
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	User := domain.User{ID: id}

	User, err = h.Services.ChangePassword(User, request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = domain.UserResponse{
		ID:    User.ID,
		Name:  User.Name,
		Email: User.Email,
		Role:  User.Role,
	}
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) ChangeRole(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.ChangeRoleRequest
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	User := domain.User{ID: id}

	User, err = h.Services.ChangeRole(User, request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = domain.UserResponse{
		ID:    User.ID,
		Name:  User.Name,
		Email: User.Email,
		Role:  User.Role,
	}
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
