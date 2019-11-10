package console

import (
	"fmt"
	"os"
	"time"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/handler"
	"github.com/impal-lms/lms-backend/repository"
	"github.com/impal-lms/lms-backend/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	middleware "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
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
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	orm := repository.NewRepository(db)
	auth := authentication.NewAuthentication()
	lms := services.NewServices(orm, auth)

	h := handler.NewHandler(lms)

	e := echo.New()

	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	e.Use(middleware.Logger())
	log.Info("Logger enabled!!")

	e.GET("/", h.HelloWorld)
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)

	user := e.Group("/user")
	user.GET("", h.GetAllUser)
	user.POST("", h.CreateUser)
	user.GET("/:id", h.GetUserByID)
	user.PUT("/:id", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUserById)
	user.PUT("/:id/password", h.ChangePassword)
	user.PUT("/:id/role", h.ChangeRole)

	classroom := e.Group("/classroom")
	classroom.GET("", h.GetAllClassroom)
	classroom.POST("", h.CreateClassroom)
	classroom.GET("/:id", h.GetClassroomByID)
	classroom.PUT("/:id", h.UpdateClassroom)
	classroom.DELETE("/:id", h.DeleteClassroomById)
	classroom.PUT("/:id/add-student", h.AddStudentToClassroom)
	classroom.PUT("/:id/delete-student", h.DeleteStudentFromClassroom)
	classroom.PUT("/:id/add-room", h.AddRoomToClassroom)
	classroom.PUT("/:id/delete-room", h.DeleteRoomFromClassroom)
	classroom.GET("/user/:user_id", h.GetAllClassroomOfUser)

	material := e.Group("/material")
	material.GET("", h.GetAllMaterial)
	material.POST("", h.CreateMaterial)
	material.GET("/:id", h.GetMaterialByID)
	material.PUT("/:id", h.UpdateMaterial)
	material.DELETE("/:id", h.DeleteMaterialById)

	task := e.Group("/task")
	task.GET("", h.GetAllTask)
	task.POST("", h.CreateTask)
	task.GET("/:id", h.GetTaskByID)
	task.PUT("/:id", h.UpdateTask)
	task.DELETE("/:id", h.DeleteTaskById)

	submission := e.Group("/submission")
	submission.GET("", h.GetAllSubmission)
	submission.POST("", h.CreateSubmission)
	submission.GET("/:id", h.GetSubmissionByID)
	submission.PUT("/:id", h.UpdateSubmission)
	submission.DELETE("/:id", h.DeleteSubmissionById)

	file := e.Group("/file")
	file.POST("/upload", h.FileUpload)
	file.Static("/", "static")

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
