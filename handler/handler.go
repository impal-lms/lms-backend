package handler

import (
	"net/http"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Handler struct {
	Services services.Services
}

func NewHandler(services services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}

/*
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

	resp, err := h.userService.Login(req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
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

func (h Handler) FileUpload(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	name := strings.TrimSuffix(file.Filename, path.Ext(file.Filename))
	ekstension := filepath.Ext(file.Filename)
	if !domain.IsValidExtension(ekstension) {
		return echo.ErrUnsupportedMediaType
	}

	src, err := file.Open()
	if err != nil {
		return err
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
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"path": "/file/" + filename,
	})
}

*/
