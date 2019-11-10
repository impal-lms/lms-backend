package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func (lms *LMS) DeleteClassroomById(id int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	return domain.Classroom{}, lms.Repository.DeleteClassroomById(Classroom.ID)
}

func (lms *LMS) AddStudentToClassroom(id int64, studentID int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	var students []int64

	if studentID == 0 {
		return domain.Classroom{}, http.ErrBodyNotAllowed
	}

	if Classroom.StudentIDs == nil {
		students = append(students, studentID)
	}

	students = Classroom.StudentIDs
	for i := range students {
		if students[i] == studentID {
			return domain.Classroom{}, errors.Wrap(err, "student already exists")
		}
	}

	students = append(students, studentID)
	Classroom.StudentIDs = students

	return Classroom, lms.Repository.UpdateClassroom(Classroom)
}

func (lms *LMS) DeleteStudentFromClassroom(id int64, studentID int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	if Classroom.StudentIDs == nil {
		return domain.Classroom{}, errors.Wrap(err, "student list is empty")
	}

	var found bool
	var idx int

	students := Classroom.StudentIDs
	for i := range students {
		if students[i] == studentID {
			idx = i
			found = true
		}
	}

	if !found {
		return domain.Classroom{}, errors.Wrap(err, "student id does not exist")
	}

	students = append(students[:idx], students[idx+1:]...)
	Classroom.StudentIDs = students

	return Classroom, lms.Repository.UpdateClassroom(Classroom)
}

func (lms *LMS) AddRoomToClassroom(id int64, roomID int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	var rooms []int64

	if roomID == 0 {
		return domain.Classroom{}, http.ErrBodyNotAllowed
	}

	if Classroom.RoomIDs == nil {
		rooms = append(rooms, roomID)
	}

	rooms = Classroom.RoomIDs
	for i := range rooms {
		if rooms[i] == roomID {
			return domain.Classroom{}, errors.Wrap(err, "room already exists")
		}
	}

	rooms = append(rooms, roomID)
	Classroom.RoomIDs = rooms

	return Classroom, lms.Repository.UpdateClassroom(Classroom)
}

func (lms *LMS) DeleteRoomFromClassroom(id int64, roomID int64) (domain.Classroom, error) {
	Classroom, err := lms.Repository.GetClassroomByID(id)
	if err != nil {
		return domain.Classroom{}, errors.Wrap(err, "Classroom does not exist")
	}

	if Classroom.RoomIDs == nil {
		return domain.Classroom{}, errors.Wrap(err, "room list is empty")
	}

	var found bool
	var idx int

	roooms := Classroom.RoomIDs
	for i := range roooms {
		if roooms[i] == roomID {
			idx = i
			found = true
		}
	}

	if !found {
		return domain.Classroom{}, errors.Wrap(err, "room id does not exist")
	}

	roooms = append(roooms[:idx], roooms[idx+1:]...)
	Classroom.RoomIDs = roooms

	return Classroom, lms.Repository.UpdateClassroom(Classroom)
}

func (lms *LMS) GetAllClassroomOfUser(userID int64) ([]domain.ClassroomOfUser, error) {
	user, err := lms.GetUserByID(userID)
	if err != nil {
		return []domain.ClassroomOfUser{}, errors.Wrap(err, "user does not exist")
	}

	var classesOfUser []domain.ClassroomOfUser

	classes, err := lms.GetAllClassroom()
	if err != nil {
		return []domain.ClassroomOfUser{}, err
	}

	if user.Role == domain.Teacher {
		for _, class := range classes {
			if class.TeacherID == user.ID {
				classesOfUser = append(classesOfUser, domain.ClassroomOfUser{
					ID:        class.ID,
					Label:     class.Label,
					TeacherID: class.TeacherID,
				})
			}
		}
	} else if user.Role == domain.Student {
		logrus.Println(classes)
		for _, class := range classes {
			students := class.StudentIDs
			for _, student := range students {
				if user.ID == student {
					classesOfUser = append(classesOfUser, domain.ClassroomOfUser{
						ID:        class.ID,
						Label:     class.Label,
						TeacherID: class.TeacherID,
					})
				}
			}
		}
	}

	return classesOfUser, nil
}
