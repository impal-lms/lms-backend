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

	GetAllMaterial() ([]domain.Material, error)
	CreateMaterial(Material domain.Material) (domain.Material, int, error)
	GetMaterialByID(id int64) (domain.Material, error)
	UpdateMaterial(Material domain.Material) (domain.Material, error)
	DeleteMaterialById(id int64) (domain.Material, error)

	GetAllTask() ([]domain.Task, error)
	CreateTask(Task domain.Task) (domain.Task, int, error)
	GetTaskByID(id int64) (domain.Task, error)
	UpdateTask(Task domain.Task) (domain.Task, error)
	DeleteTaskById(id int64) (domain.Task, error)

	GetAllSubmission() ([]domain.Submission, error)
	CreateSubmission(Submission domain.Submission) (domain.Submission, int, error)
	GetSubmissionByID(id int64) (domain.Submission, error)
	UpdateSubmission(Submission domain.Submission) (domain.Submission, error)
	DeleteSubmissionById(id int64) (domain.Submission, error)
}
