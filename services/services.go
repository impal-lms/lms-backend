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
	DeleteUserById(id int64) (domain.User, error)
	ChangePassword(user domain.User, req domain.ChangePasswordRequest) (domain.User, error)
	ChangeRole(user domain.User, req domain.ChangeRoleRequest) (domain.User, error)

	GetAllClassroom() ([]domain.Classroom, error)
	CreateClassroom(Classroom domain.Classroom) (domain.Classroom, int, error)
	GetClassroomByID(id int64) (domain.Classroom, error)
	UpdateClassroom(Classroom domain.Classroom) (domain.Classroom, error)
	DeleteClassroomById(id int64) (domain.Classroom, error)
	AddStudentToClassroom(id int64, studentID int64) (domain.Classroom, error)
	DeleteStudentFromClassroom(id int64, studentID int64) (domain.Classroom, error)
	AddRoomToClassroom(id int64, roomID int64) (domain.Classroom, error)
	DeleteRoomFromClassroom(id int64, roomID int64) (domain.Classroom, error)
	GetAllClassroomOfUser(userID int64) ([]domain.ClassroomOfUser, error)

	GetAllMaterial(classroomID int64) ([]domain.Material, error)
	CreateMaterial(Material domain.Material) (domain.Material, int, error)
	GetMaterialByID(id int64) (domain.Material, error)
	UpdateMaterial(Material domain.Material) (domain.Material, error)
	DeleteMaterialById(id int64) (domain.Material, error)

	GetAllTask(classroomID int64) ([]domain.Task, error)
	CreateTask(Task domain.Task) (domain.Task, int, error)
	GetTaskByID(id int64) (domain.Task, error)
	UpdateTask(Task domain.Task) (domain.Task, error)
	DeleteTaskById(id int64) (domain.Task, error)

	GetAllSubmission(classroomID int64) ([]domain.Submission, error)
	CreateSubmission(Submission domain.Submission) (domain.Submission, int, error)
	GetSubmissionByID(id int64) (domain.Submission, error)
	UpdateSubmission(Submission domain.Submission) (domain.Submission, error)
	DeleteSubmissionById(id int64) (domain.Submission, error)
}
