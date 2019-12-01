package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllChatRoom(ctx echo.Context) error {
	var response Response

	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	ChatRooms, err := h.Services.GetAllChatRoom(userID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = ChatRooms
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateChatRoom(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.ChatRoom
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	ChatRoom, code, err := h.Services.CreateChatRoom(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = ChatRoom
	response.Status = code

	return ctx.JSON(code, response)
}

func (h *Handler) GetChatRoomByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	ChatRoom, err := h.Services.GetChatRoomByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = ChatRoom
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateChatRoom(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.ChatRoom
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	request.ID = id

	ChatRoom, err := h.Services.UpdateChatRoom(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = ChatRoom
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
func (h *Handler) DeleteChatRoomByID(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeleteChatRoomByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllMessage(ctx echo.Context) error {
	var response Response

	chatRoomID, err := strconv.ParseInt(ctx.Param("chat_room_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Messages, err := h.Services.GetAllMessage(chatRoomID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Messages
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateMessage(ctx echo.Context) error {
	var response Response

	body := ctx.Request().Body
	var request domain.Message
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Message, code, err := h.Services.CreateMessage(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Message
	response.Status = code

	return ctx.JSON(code, response)
}

func (h *Handler) GetMessageByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Message, err := h.Services.GetMessageByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Message
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateMessage(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	body := ctx.Request().Body
	var request domain.Message
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	request.ID = id

	Message, err := h.Services.UpdateMessage(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = Message
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
func (h *Handler) DeleteMessageByID(ctx echo.Context) error {
	var response Response

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	_, err = h.Services.DeleteMessageByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Data = nil
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
