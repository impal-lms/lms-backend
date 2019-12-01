package services

import (
	"strings"
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllClassroom() ([]domain.Classroom, error) {
	return lms.Repository.GetAllClassroom()
}

func (lms *LMS) CreateClassroom(Classroom domain.Classroom) (domain.Classroom, int, error) {
	Classroom.ID = time.Now().UnixNano()
	Classroom.Label = strings.ToUpper(Classroom.Label)

	Classroom, err := lms.Repository.CreateClassroom(Classroom)
	if err != nil {
		return domain.Classroom{}, 400, errors.Wrap(err, "create Classroom error")
	}

	return Classroom, 201, nil
}

func (lms *LMS) GetClassroomByID(id int64) (domain.Classroom, error) {
	return lms.Repository.GetClassroomByID(id)
}

func (lms *LMS) UpdateClassroom(Classroom domain.Classroom) (domain.Classroom, error) {
	_, err := lms.Repository.GetClassroomByID(Classroom.ID)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	if Classroom.Label != "" {
		Classroom.Label = strings.ToUpper(Classroom.Label)
	}

	return Classroom, lms.Repository.UpdateClassroom(Classroom)
}

func (lms *LMS) DeleteClassroomByID(id int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	return domain.Classroom{}, lms.Repository.DeleteClassroomByID(Classroom.ID)
}

func (lms *LMS) GetAllStudentClassroom(studentID, classroomID int64) ([]domain.StudentClassroom, error) {
	return lms.Repository.GetAllStudentClassroom(studentID, classroomID)
}

func (lms *LMS) CreateStudentClassroom(StudentClassroom domain.StudentClassroom) (domain.StudentClassroom, int, error) {
	StudentClassroom.ID = time.Now().UnixNano()

	StudentClassroom, err := lms.Repository.CreateStudentClassroom(StudentClassroom)
	if err != nil {
		return domain.StudentClassroom{}, 400, errors.Wrap(err, "create Classroom error")
	}

	return StudentClassroom, 201, nil
}

func (lms *LMS) GetStudentClassroomByID(id int64) (domain.StudentClassroom, error) {
	return lms.Repository.GetStudentClassroomByID(id)
}

func (lms *LMS) UpdateStudentClassroom(StudentClassroom domain.StudentClassroom) (domain.StudentClassroom, error) {
	_, err := lms.Repository.GetStudentClassroomByID(StudentClassroom.ID)
	if err != nil {
		return domain.StudentClassroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	return StudentClassroom, lms.Repository.UpdateStudentClassroom(StudentClassroom)
}

func (lms *LMS) DeleteStudentClassroomByID(id int64) (domain.StudentClassroom, error) {
	StudentClassroom, err := lms.Repository.GetStudentClassroomByID(id)
	if err != nil {
		return domain.StudentClassroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	return domain.StudentClassroom{}, lms.Repository.DeleteStudentClassroomByID(StudentClassroom.ID)
}
