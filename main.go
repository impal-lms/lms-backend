package main

import (
	"github.com/impal-lms/lms-backend/handler"
	"github.com/labstack/echo"
)

func main() {
	h := handler.New()

	e := echo.New()
	e.GET("/", h.HelloWorld)
	e.POST("/login", h.Login)

	user := e.Group("/user")
	user.POST("/create", h.CreateUser)
	user.GET("/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(":8080"))
}
