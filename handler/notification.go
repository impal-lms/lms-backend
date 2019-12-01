package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllNotification(ctx echo.Context) error {
	var response Response

	studentID, err := strconv.ParseInt(ctx.Param("student_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Notifications, err := h.Services.GetAllNotification(studentID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Notifications
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateNotification(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.Notification
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Notification, code, err := h.Services.CreateNotification(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Notification
	response.Status = code

	return ctx.JSON(code, response)
}

func (h *Handler) GetNotificationByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Notification, err := h.Services.GetNotificationByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Notification
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateNotification(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.Notification
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	request.ID = id

	Notification, err := h.Services.UpdateNotification(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Notification
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
func (h *Handler) DeleteNotificationByID(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeleteNotificationByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
