package repository

import "github.com/impal-lms/lms-backend/domain"

type UserRepository interface {
	GetUserByID(id int64) (*domain.User, error)
	Authenticate(email string, password string) (*domain.User, error)
	Save(user domain.User) error
}

type ClassroomRepository interface {
	Save(classroom domain.Classroom) error
	GetClassroomByID(id int64) (*domain.Classroom, error)
	GetClassroomByStudentID(id int64) int64
	DeleteClassroomById(id int64) error
}

type FileRepository interface {
	SaveTask(task domain.Task) error
	SaveSubmission(submission domain.Submission) error
	GetTaskByID(id int64) (*domain.Task, error)
}
