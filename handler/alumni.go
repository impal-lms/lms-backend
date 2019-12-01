package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllAlumni(ctx echo.Context) error {
	var response Response

	Alumnis, err := h.Services.GetAllAlumni()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Alumnis
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateAlumni(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.Alumni
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Alumni, code, err := h.Services.CreateAlumni(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Alumni
	response.Status = code

	return ctx.JSON(code, response)
}

func (h *Handler) GetAlumniByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Alumni, err := h.Services.GetAlumniByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Alumni
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateAlumni(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.Alumni
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	request.ID = id

	Alumni, err := h.Services.UpdateAlumni(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Alumni
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
func (h *Handler) DeleteAlumniByID(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeleteAlumniByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
