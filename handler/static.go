package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FileUpload(ctx echo.Context) error {
	var response Response

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)

	}

	name := strings.TrimSuffix(file.Filename, path.Ext(file.Filename))
	ekstension := filepath.Ext(file.Filename)
	if !services.IsValidExtension(ekstension) {
		response.Data = err.Error()
		response.Status = 415
		return ctx.JSON(415, response)
	}

	src, err := file.Open()
	if err != nil {
		response.Data = err.Error()
		response.Status = 500
		return ctx.JSON(500, response)
	}

	defer src.Close()

	path := "static/"
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}

	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), name, ekstension)
	dst, err := os.Create(path + filename)
	if err != nil {
		response.Data = err.Error()
		response.Status = 500
		return ctx.JSON(500, response)
	}

	if _, err = io.Copy(dst, src); err != nil {
		response.Data = err.Error()
		response.Status = 500
		return ctx.JSON(500, response)
	}

	response.Data = map[string]string{
		"path": "/file/" + filename,
	}
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
