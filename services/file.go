package services

import (
	"strings"
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func IsValidExtension(ext string) bool {
	for _, all := range domain.AllowedExtension {
		if ext == all {
			return true
		}
	}
	return false
}

func (lms *LMS) GetAllMaterial(classroomID int64) ([]domain.Material, error) {
	return lms.Repository.GetAllMaterial(classroomID)
}

func (lms *LMS) CreateMaterial(Material domain.Material) (domain.Material, int, error) {
	Material.ID = time.Now().UnixNano()
	if Material.Title != "" {
		Material.Title = strings.ToUpper(Material.Title)
	}

	Material, err := lms.Repository.CreateMaterial(Material)
	if err != nil {
		return domain.Material{}, 400, errors.Wrap(err, "create material error")
	}

	Students, err := lms.GetAllStudentClassroom(0, Material.ClassroomID)
	if err != nil {
		return domain.Material{}, 400, errors.Wrap(err, "failed to create notification")
	}

	SendAt := time.Now()
	SendAtIndex := SendAt.UnixNano()

	for i := range Students {
		var notification domain.Notification
		notification.ReceiverID = Students[i].StudentID
		notification.SendAt = SendAt
		notification.SendAtIndex = SendAtIndex
		notification.Content = "New Material Has Been Posted"

		notification, _, err = lms.CreateNotification(notification)
		if err != nil {
			return domain.Material{}, 400, errors.Wrap(err, "failed to create notification")
		}

	}

	return Material, 201, nil
}

func (lms *LMS) GetMaterialByID(id int64) (domain.Material, error) {
	return lms.Repository.GetMaterialByID(id)
}

func (lms *LMS) UpdateMaterial(Material domain.Material) (domain.Material, error) {
	_, err := lms.Repository.GetMaterialByID(Material.ID)
	if err != nil {
		return domain.Material{}, errors.Wrap(err, "material does not exist")
	}

	if Material.Title != "" {
		Material.Title = strings.ToUpper(Material.Title)
	}

	return Material, lms.Repository.UpdateMaterial(Material)
}

func (lms *LMS) DeleteMaterialByID(id int64) (domain.Material, error) {
	Material, err := lms.Repository.GetMaterialByID(id)
	if err != nil {
		return domain.Material{}, errors.Wrap(err, "material does not exist")
	}

	return domain.Material{}, lms.Repository.DeleteMaterialByID(Material.ID)
}

func (lms *LMS) GetAllTask(classroomID int64) ([]domain.Task, error) {
	return lms.Repository.GetAllTask(classroomID)
}

func (lms *LMS) CreateTask(Task domain.Task) (domain.Task, int, error) {
	Task.ID = time.Now().UnixNano()
	if Task.Title != "" {
		Task.Title = strings.ToUpper(Task.Title)
	}

	Task, err := lms.Repository.CreateTask(Task)
	if err != nil {
		return domain.Task{}, 400, errors.Wrap(err, "create task error")
	}

	Students, err := lms.GetAllStudentClassroom(0, Task.ClassroomID)
	if err != nil {
		return domain.Task{}, 400, errors.Wrap(err, "failed to create notification")
	}

	SendAt := time.Now()
	SendAtIndex := SendAt.UnixNano()

	for i := range Students {
		var notification domain.Notification
		notification.ReceiverID = Students[i].StudentID
		notification.SendAt = SendAt
		notification.SendAtIndex = SendAtIndex
		notification.Content = "New Task Has Been Posted"

		notification, _, err = lms.CreateNotification(notification)
		if err != nil {
			return domain.Task{}, 400, errors.Wrap(err, "failed to create notification")
		}

	}

	return Task, 201, nil
}

func (lms *LMS) GetTaskByID(id int64) (domain.Task, error) {
	return lms.Repository.GetTaskByID(id)
}

func (lms *LMS) UpdateTask(Task domain.Task) (domain.Task, error) {
	_, err := lms.Repository.GetTaskByID(Task.ID)
	if err != nil {
		return domain.Task{}, errors.Wrap(err, "task does not exist")
	}

	if Task.Title != "" {
		Task.Title = strings.ToUpper(Task.Title)
	}

	return Task, lms.Repository.UpdateTask(Task)
}

func (lms *LMS) DeleteTaskByID(id int64) (domain.Task, error) {
	Task, err := lms.Repository.GetTaskByID(id)
	if err != nil {
		return domain.Task{}, errors.Wrap(err, "task does not exist")
	}

	return domain.Task{}, lms.Repository.DeleteTaskByID(Task.ID)
}

func (lms *LMS) GetAllSubmission(studentID, classroomID int64) ([]domain.Submission, error) {
	return lms.Repository.GetAllSubmission(studentID, classroomID)
}

func (lms *LMS) CreateSubmission(Submission domain.Submission) (domain.Submission, int, error) {
	Submission.ID = time.Now().UnixNano()
	if Submission.Title != "" {
		Submission.Title = strings.ToUpper(Submission.Title)
	}

	Submission, err := lms.Repository.CreateSubmission(Submission)
	if err != nil {
		return domain.Submission{}, 400, errors.Wrap(err, "create submission error")
	}

	return Submission, 201, nil
}

func (lms *LMS) GetSubmissionByID(id int64) (domain.Submission, error) {
	return lms.Repository.GetSubmissionByID(id)
}

func (lms *LMS) UpdateSubmission(Submission domain.Submission) (domain.Submission, error) {
	_, err := lms.Repository.GetSubmissionByID(Submission.ID)
	if err != nil {
		return domain.Submission{}, errors.Wrap(err, "submission does not exist")
	}

	if Submission.Title != "" {
		Submission.Title = strings.ToUpper(Submission.Title)
	}

	return Submission, lms.Repository.UpdateSubmission(Submission)
}

func (lms *LMS) DeleteSubmissionByID(id int64) (domain.Submission, error) {
	Submission, err := lms.Repository.GetSubmissionByID(id)
	if err != nil {
		return domain.Submission{}, errors.Wrap(err, "submission does not exist")
	}

	return domain.Submission{}, lms.Repository.DeleteSubmissionByID(Submission.ID)
}
