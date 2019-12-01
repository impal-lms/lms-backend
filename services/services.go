package services

import "github.com/impal-lms/lms-backend/domain"

type Services interface {
	Authenticate(email, password string) (*domain.User, error)
	Login(req domain.LoginRequest) (domain.LoginResponse, error)

	GetAllUser() ([]domain.User, error)
	CreateUser(user domain.User) (domain.User, int, error)
	GetUserByID(id int64) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUserByID(id int64) (domain.User, error)
	ChangePassword(user domain.User, req domain.ChangePasswordRequest) (domain.User, error)
	ChangeRole(user domain.User, req domain.ChangeRoleRequest) (domain.User, error)

	GetAllClassroom() ([]domain.Classroom, error)
	CreateClassroom(Classroom domain.Classroom) (domain.Classroom, int, error)
	GetClassroomByID(id int64) (domain.Classroom, error)
	UpdateClassroom(Classroom domain.Classroom) (domain.Classroom, error)
	DeleteClassroomByID(id int64) (domain.Classroom, error)

	GetAllStudentClassroom(studentID, classroomID int64) ([]domain.StudentClassroom, error)
	CreateStudentClassroom(StudentClassroom domain.StudentClassroom) (domain.StudentClassroom, int, error)
	GetStudentClassroomByID(id int64) (domain.StudentClassroom, error)
	UpdateStudentClassroom(StudentClassroom domain.StudentClassroom) (domain.StudentClassroom, error)
	DeleteStudentClassroomByID(id int64) (domain.StudentClassroom, error)

	GetAllMaterial(classroomID int64) ([]domain.Material, error)
	CreateMaterial(Material domain.Material) (domain.Material, int, error)
	GetMaterialByID(id int64) (domain.Material, error)
	UpdateMaterial(Material domain.Material) (domain.Material, error)
	DeleteMaterialByID(id int64) (domain.Material, error)

	GetAllTask(classroomID int64) ([]domain.Task, error)
	CreateTask(Task domain.Task) (domain.Task, int, error)
	GetTaskByID(id int64) (domain.Task, error)
	UpdateTask(Task domain.Task) (domain.Task, error)
	DeleteTaskByID(id int64) (domain.Task, error)

	GetAllSubmission(studentID, classroomID int64) ([]domain.Submission, error)
	CreateSubmission(Submission domain.Submission) (domain.Submission, int, error)
	GetSubmissionByID(id int64) (domain.Submission, error)
	UpdateSubmission(Submission domain.Submission) (domain.Submission, error)
	DeleteSubmissionByID(id int64) (domain.Submission, error)

	GetAllAlumni() ([]domain.Alumni, error)
	CreateAlumni(Alumni domain.Alumni) (domain.Alumni, int, error)
	GetAlumniByID(id int64) (domain.Alumni, error)
	UpdateAlumni(Alumni domain.Alumni) (domain.Alumni, error)
	DeleteAlumniByID(id int64) (domain.Alumni, error)

	GetAllPresence(studentID int64) ([]domain.Presence, error)
	CreatePresence(Presence domain.Presence) (domain.Presence, int, error)
	GetPresenceByID(id int64) (domain.Presence, error)
	UpdatePresence(Presence domain.Presence) (domain.Presence, error)
	DeletePresenceByID(id int64) (domain.Presence, error)

	GetAllChatRoom(userID int64) ([]domain.ChatRoom, error)
	GetChatRoomByID(ChatRoomID int64) (domain.ChatRoom, error)
	CreateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, int, error)
	UpdateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, error)
	DeleteChatRoomByID(ChatRoomID int64) (domain.ChatRoom, error)

	GetAllMessage(ChatRoomID int64) ([]domain.Message, error)
	GetMessageByID(messageID int64) (domain.Message, error)
	CreateMessage(message domain.Message) (domain.Message, int, error)
	UpdateMessage(message domain.Message) (domain.Message, error)
	DeleteMessageByID(messageID int64) (domain.Message, error)

	GetAllNotification(userID int64) ([]domain.Notification, error)
	CreateNotification(notification domain.Notification) (domain.Notification, int, error)
	GetNotificationByID(ID int64) (domain.Notification, error)
	UpdateNotification(notification domain.Notification) (domain.Notification, error)
	DeleteNotificationByID(ID int64) (domain.Notification, error)
}
