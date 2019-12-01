package repository

import "github.com/impal-lms/lms-backend/domain"

type Repository interface {
	GetAllUser() ([]domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	GetUserByID(id int64) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUserByID(id int64) error

	GetAllClassroom() ([]domain.Classroom, error)
	CreateClassroom(Classroom domain.Classroom) (domain.Classroom, error)
	GetClassroomByID(id int64) (domain.Classroom, error)
	UpdateClassroom(Classroom domain.Classroom) error
	DeleteClassroomByID(id int64) error

	GetAllStudentClassroom(studentID, classroomID int64) ([]domain.StudentClassroom, error)
	CreateStudentClassroom(StudentClassroom domain.StudentClassroom) (domain.StudentClassroom, error)
	GetStudentClassroomByID(id int64) (domain.StudentClassroom, error)
	UpdateStudentClassroom(StudentClassroom domain.StudentClassroom) error
	DeleteStudentClassroomByID(id int64) error

	GetAllMaterial(classroomID int64) ([]domain.Material, error)
	CreateMaterial(Material domain.Material) (domain.Material, error)
	GetMaterialByID(id int64) (domain.Material, error)
	UpdateMaterial(Material domain.Material) error
	DeleteMaterialByID(id int64) error

	GetAllTask(classroomID int64) ([]domain.Task, error)
	CreateTask(Task domain.Task) (domain.Task, error)
	GetTaskByID(id int64) (domain.Task, error)
	UpdateTask(Task domain.Task) error
	DeleteTaskByID(id int64) error

	GetAllSubmission(studentID, classroomID int64) ([]domain.Submission, error)
	CreateSubmission(Submission domain.Submission) (domain.Submission, error)
	GetSubmissionByID(id int64) (domain.Submission, error)
	UpdateSubmission(Submission domain.Submission) error
	DeleteSubmissionByID(id int64) error

	GetAllAlumni() ([]domain.Alumni, error)
	CreateAlumni(alumni domain.Alumni) (domain.Alumni, error)
	GetAlumniByID(id int64) (domain.Alumni, error)
	UpdateAlumni(alumni domain.Alumni) error
	DeleteAlumniByID(id int64) error

	GetAllPresence(studentID int64) ([]domain.Presence, error)
	CreatePresence(Presence domain.Presence) (domain.Presence, error)
	GetPresenceByID(id int64) (domain.Presence, error)
	UpdatePresence(Presence domain.Presence) error
	DeletePresenceByID(id int64) error

	GetAllChatRoom(userID int64) ([]domain.ChatRoom, error)
	CreateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, error)
	GetChatRoomByID(id int64) (domain.ChatRoom, error)
	UpdateChatRoom(ChatRoom domain.ChatRoom) error
	DeleteChatRoomByID(id int64) error
	GetChatRoomBetweenUsers(studentID, teacherID int64) (domain.ChatRoom, error)

	GetAllMessage(ChatRoomID int64) ([]domain.Message, error)
	CreateMessage(Message domain.Message) (domain.Message, error)
	GetMessageByID(id int64) (domain.Message, error)
	UpdateMessage(Message domain.Message) error
	DeleteMessageByID(id int64) error

	GetAllNotification(userID int64) ([]domain.Notification, error)
	CreateNotification(Notification domain.Notification) (domain.Notification, error)
	GetNotificationByID(id int64) (domain.Notification, error)
	UpdateNotification(Notification domain.Notification) error
	DeleteNotificationByID(id int64) error
}
