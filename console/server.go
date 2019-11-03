package console

import (
	"os"

	"github.com/impal-lms/lms-backend/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
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
