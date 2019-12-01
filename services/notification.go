package services

import (
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllNotification(userID int64) ([]domain.Notification, error) {
	return lms.Repository.GetAllNotification(userID)
}

func (lms *LMS) CreateNotification(notification domain.Notification) (domain.Notification, int, error) {
	notification.ID = time.Now().UnixNano()

	notification, err := lms.Repository.CreateNotification(notification)
	if err != nil {
		return domain.Notification{}, 400, errors.Wrap(err, "Create notification error")
	}

	return notification, 201, nil
}

func (lms *LMS) GetNotificationByID(ID int64) (domain.Notification, error) {
	notification, err := lms.Repository.GetNotificationByID(ID)
	if err != nil {
		return domain.Notification{}, err
	}

	return notification, nil
}

func (lms *LMS) UpdateNotification(notification domain.Notification) (domain.Notification, error) {
	_, err := lms.Repository.GetNotificationByID(notification.ID)
	if err != nil {
		return domain.Notification{}, errors.Wrap(err, "Notification does not exist")
	}

	return notification, lms.Repository.UpdateNotification(notification)
}

func (lms *LMS) DeleteNotificationByID(ID int64) (domain.Notification, error) {
	notification, err := lms.Repository.GetNotificationByID(ID)
	if err != nil {
		return domain.Notification{}, errors.Wrap(err, "Notification does not exist")
	}

	return notification, lms.Repository.DeleteNotificationByID(notification.ID)
}
