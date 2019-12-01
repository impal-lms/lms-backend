package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllPresence(ctx echo.Context) error {
	var response Response

	studentID, err := strconv.ParseInt(ctx.Param("student_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Presences, err := h.Services.GetAllPresence(studentID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Presences
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreatePresence(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.Presence
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Presence, code, err := h.Services.CreatePresence(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Presence
	response.Status = code

	return ctx.JSON(code, response)
}

func (h *Handler) GetPresenceByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Presence, err := h.Services.GetPresenceByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Presence
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdatePresence(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.Presence
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	request.ID = id

	Presence, err := h.Services.UpdatePresence(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Presence
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
func (h *Handler) DeletePresenceByID(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeletePresenceByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
