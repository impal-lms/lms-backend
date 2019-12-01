package main

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
)

func main() {
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
	user.DELETE("/:id", h.DeleteUserByID)
	user.PUT("/:id/password", h.ChangePassword)
	user.PUT("/:id/role", h.ChangeRole)

	classroom := e.Group("/classroom")
	classroom.GET("", h.GetAllClassroom)
	classroom.POST("", h.CreateClassroom)
	classroom.GET("/:id", h.GetClassroomByID)
	classroom.PUT("/:id", h.UpdateClassroom)
	classroom.DELETE("/:id", h.DeleteClassroomByID)

	studentClassroom := e.Group("/student-classroom")
	studentClassroom.GET("", h.GetAllStudentClassroom)
	studentClassroom.POST("", h.CreateStudentClassroom)
	studentClassroom.GET("/:id", h.GetStudentClassroomByID)
	studentClassroom.PUT("/:id", h.UpdateStudentClassroom)
	studentClassroom.DELETE("/:id", h.DeleteStudentClassroomByID)

	material := e.Group("/material")
	material.GET("", h.GetAllMaterial)
	material.POST("", h.CreateMaterial)
	material.GET("/:id", h.GetMaterialByID)
	material.PUT("/:id", h.UpdateMaterial)
	material.DELETE("/:id", h.DeleteMaterialByID)

	task := e.Group("/task")
	task.GET("", h.GetAllTask)
	task.POST("", h.CreateTask)
	task.GET("/:id", h.GetTaskByID)
	task.PUT("/:id", h.UpdateTask)
	task.DELETE("/:id", h.DeleteTaskByID)

	submission := e.Group("/submission")
	submission.GET("", h.GetAllSubmission)
	submission.POST("", h.CreateSubmission)
	submission.GET("/:id", h.GetSubmissionByID)
	submission.PUT("/:id", h.UpdateSubmission)
	submission.DELETE("/:id", h.DeleteSubmissionByID)

	alumni := e.Group("/alumni")
	alumni.GET("", h.GetAllAlumni)
	alumni.POST("", h.CreateAlumni)
	alumni.GET("/:id", h.GetAlumniByID)
	alumni.PUT("/:id", h.UpdateAlumni)
	alumni.DELETE("/:id", h.DeleteAlumniByID)

	presence := e.Group("/presence")
	presence.GET("", h.GetAllPresence)
	presence.POST("", h.CreatePresence)
	presence.GET("/:id", h.GetPresenceByID)
	presence.PUT("/:id", h.UpdatePresence)
	presence.DELETE("/:id", h.DeletePresenceByID)

	chatRoom := e.Group("/chat-room")
	chatRoom.GET("", h.GetAllChatRoom)
	chatRoom.POST("", h.CreateChatRoom)
	chatRoom.GET("/:id", h.GetChatRoomByID)
	chatRoom.PUT("/:id", h.UpdateChatRoom)
	chatRoom.DELETE("/:id", h.DeleteChatRoomByID)

	message := e.Group("/message")
	message.GET("", h.GetAllMessage)
	message.POST("", h.CreateMessage)
	message.GET("/:id", h.GetMessageByID)
	message.PUT("/:id", h.UpdateMessage)
	message.DELETE("/:id", h.DeleteMessageByID)

	notification := e.Group("/notification")
	notification.GET("", h.GetAllNotification)
	notification.POST("", h.CreateNotification)
	notification.GET("/:id", h.GetNotificationByID)
	notification.PUT("/:id", h.UpdateNotification)
	notification.DELETE("/:id", h.DeleteNotificationByID)

	file := e.Group("/file")
	file.POST("/upload", h.FileUpload)
	file.Static("/", "static")

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
