package repository

import "github.com/impal-lms/lms-backend/domain"

type Repository interface {
	GetAllUser() ([]domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	GetUserByID(id int64) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUserById(id int64) error

	GetAllMaterial() ([]domain.Material, error)
	CreateMaterial(Material domain.Material) (domain.Material, error)
	GetMaterialByID(id int64) (domain.Material, error)
	UpdateMaterial(Material domain.Material) error
	DeleteMaterialById(id int64) error

	GetAllTask() ([]domain.Task, error)
	CreateTask(Task domain.Task) (domain.Task, error)
	GetTaskByID(id int64) (domain.Task, error)
	UpdateTask(Task domain.Task) error
	DeleteTaskById(id int64) error

	GetAllSubmission() ([]domain.Submission, error)
	CreateSubmission(Submission domain.Submission) (domain.Submission, error)
	GetSubmissionByID(id int64) (domain.Submission, error)
	UpdateSubmission(Submission domain.Submission) error
	DeleteSubmissionById(id int64) error
}
