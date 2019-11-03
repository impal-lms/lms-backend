package console

import (
	"github.com/impal-lms/lms-backend/handler"
	"github.com/impal-lms/lms-backend/helper"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(_ *cobra.Command, _ []string) {
	port := helper.GetEnv("PORT", ":8000").(string)
	h := handler.New()

	e := echo.New()
	e.GET("/", h.HelloWorld)
	e.POST("/login", h.Login)

	user := e.Group("/user")
	user.POST("/create", h.CreateUser)
	user.GET("/:id", h.GetUserById)

	e.Logger.Fatal(e.Start(port))
}
