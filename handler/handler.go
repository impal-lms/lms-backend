package handler

import "github.com/labstack/echo"
import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}
