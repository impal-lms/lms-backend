package main

import (
	"os"

	"github.com/impal-lms/lms-backend/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	baseUrl := ":8000"
	if err != nil {
		logrus.Warn("Error loading .env file")
	}

	baseUrl = os.Getenv("BASE_URL")

	h := handler.New()

	e := echo.New()
	e.GET("/", h.HelloWorld)
	e.POST("/login", h.Login)

	user := e.Group("/user")
	user.POST("/create", h.CreateUser)
	user.GET("/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(baseUrl))
}
